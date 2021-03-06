package runtime

import (
	"context"
	"fmt"
	"time"

	builderinterface "github.com/level11consulting/orbitalci/build/builder/interface"
	stringbuilder "github.com/level11consulting/orbitalci/build/helpers/stringbuilder/accountrepo"
	"github.com/level11consulting/orbitalci/build/buildmonitor"
	"github.com/level11consulting/orbitalci/models"
	"github.com/level11consulting/orbitalci/models/pb"
	"github.com/level11consulting/orbitalci/storage"
	"github.com/prometheus/client_golang/prometheus"
	ocelog "github.com/shankj3/go-til/log"
)

var (
	activeBuilds = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "ocelot_active_builds",
			Help: "Number of builds currently in progress",
		},
	)
	buildDurationHist = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ocelot_build_duration_seconds",
			Help:    "Build Duration distribution",
			Buckets: []float64{1, 10, 30, 60, 120, 200},
		},
		[]string{"werker_type"},
	)
	buildCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ocelot_build_count_total",
			Help: "number of ocelot builds executed",
		},
	)
)

func init() {
	prometheus.MustRegister(activeBuilds, buildDurationHist, buildCount)
}

func startBuild() time.Time {
	time.Now()
	buildCount.Inc()
	activeBuilds.Inc()
	return time.Now()
}

func endBuild(werkType string, start time.Time) {
	activeBuilds.Dec()
	buildDurationHist.WithLabelValues(werkType).Observe(time.Since(start).Seconds())
}

// TODO: We should stream output to standard out
// watchForResults sends the *Transport object over the transport channel for stream functions to process
func (w *launcher) WatchForResults(hash string, dbId int64) {
	ocelog.Log().Debugf("adding hash ( %s ) & infochan to transport channel", hash)
	transport := &models.Transport{Hash: hash, InfoChan: w.infochan, DbId: dbId}
	w.StreamChan <- transport
}


// MakeItSo is the bread and butter of the werker. It registers the build with consul, ensures notifications, stores all the build data
//  in the OcelotStorage implementation, and runs all the stages both setup and ones defined by the user.
func (w *launcher) MakeItSo(werk *pb.WerkerTask, builder builderinterface.Builder, finish, done chan int) {
	
	ocelog.Log().Info("Starting build")
	// Gets timestamp, and increments metrics for prometheus
	startTime := startBuild()

	ocelog.Log().Debug("hash build ", werk.CheckoutHash)


	w.BuildMonitor.RegisterDoneChan(werk.CheckoutHash, done)

	// Cleanup function?
	defer w.BuildMonitor.MakeItSoDed(finish)
	// I think this is cleanup in Consul
	defer w.BuildMonitor.UnregisterDoneChan(werk.CheckoutHash)

	// TODO: Find ou
	defer func() {
		ocelog.Log().Info("calling done for nsqpb")
		done <- 1
		endBuild(w.WerkerType.String(), startTime)
	}()
	// set up notifications to be executed on build completion
	defer func() {
		if err := w.doNotifications(werk); err != nil {
			ocelog.IncludeErrField(err).Error("build notification failed!")
		}
	}()

	// create context for entire build
	ctx, cancel := context.WithCancel(context.Background())

	//send build context off, build kills are performed by calling cancel on the cancellable context
	w.BuildCtxChan <- &models.BuildContext{
		Hash:       werk.CheckoutHash,
		Context:    ctx,
		CancelFunc: cancel,
	}

	// The pattern for using context.WithCancel is to defer calling its CancelFunc
	defer cancel()

	// Get a protobuf (pb.Result) as defined by the Builder interface
	// That is, build/type/docker, or build/type/exec, or build/type/ssh etc...
	// For what it's worth, we don't use any of the parameters.
	// This is just used to customize a single string in the pg.Result 
	result := builder.Init(ctx, werk.CheckoutHash, w.infochan)


	// at the end of the build, close out any build-length connections associated with build
	defer func() {
		if err := builder.Close(); err != nil {
			ocelog.IncludeErrField(err).Error("unable to close builder connections cleanly")
		}
	}()
	if result.Status == pb.StageResultVal_FAIL {
		ocelog.Log().Error("Failed to initialize, error: " + result.Error)
		handleFailure(result, w.Store, "INIT", 0, werk.Id)
		return
	}

	// we know all the environment variables that will exist for the
	// lifetime of the container, so just add them now
	w.addGlobalEnvVars(werk, builder)
	defer func() {
		ocelog.Log().Info("closing infochan for ", werk.Id)
		close(w.infochan)
	}()

	w.WatchForResults(werk.CheckoutHash, werk.Id)

	//update consul with active build data
	consul := w.RemoteConf.GetConsul()
	// if we can't register with consul, bail, just exit out. the maintainer will soon be pausing message flow anyway
	if err := w.BuildMonitor.StartBuild(consul, werk.CheckoutHash, werk.Id); err != nil {
		return
	}

	setupStart := time.Now()
	w.BuildMonitor.Reset("setup", werk.CheckoutHash)

	// Maintainer note:
	// Is this what captures the output from docker's console?
	dockerIdChan := make(chan string)
	go w.listenForDockerUuid(dockerIdChan, werk.CheckoutHash)

	// START SETUP STAGE
	setupResult, dockerUUid := builder.Setup(ctx, w.infochan, dockerIdChan, werk, w.RemoteConf, w.ServicePort)
	defer w.BuildMonitor.Cleanup(ctx, dockerUUid, w.infochan)
	ocelog.Log().Info("finished setup")
	setupDura := time.Now().Sub(setupStart)

	if err := storeStageToDb(w.Store, werk.Id, setupResult, setupStart, setupDura.Seconds()); err != nil {
		ocelog.Log().Debug("storing failure")
		ocelog.IncludeErrField(err).Error("couldn't store build output")
		return
	}
	if setupResult.Status == pb.StageResultVal_FAIL {
		handleFailure(setupResult, w.Store, "setup", setupDura, werk.Id)
		return
	}
	// END SETUP STAGE

	// START PREFLIGHT STAGE
	// run integrations, executable download, codebase download
	if bailOut, err := w.preFlight(ctx, werk, builder); err != nil || bailOut {
		return
	}
	// END PREFLIGHT STAGE


	// run the actual stages outlined in the ocelot.yml
	fail, dura, err := w.runStages(ctx, werk, builder)
	if err != nil {
		return
	}
	// post pr comments if its relevant, send notifications
	if err := w.postFlight(ctx, werk, fail); err != nil {
		ocelog.IncludeErrField(err).Error("could not execute post flight")
		// don't return here, we still want to update the build_summary table
	}
	//update build_summary table
	if err := w.Store.UpdateSum(fail, dura.Seconds(), werk.Id); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't update summary in database")
		return
	}
	ocelog.Log().Infof("finished building id %s", werk.CheckoutHash)
}

// addGlobalEnvVars sets the global env vars on builders, these are the variables that will live for the duration of the build.
//	- `GIT_HASH`
//	- `BUILD_ID`
//	- `GIT_HASH_SHORT`
//	- `GIT_BRANCH`
//	- `WORKSPACE`
//  - `GIT_PREVIOUS_SUCCESSFUL_COMMIT`
func (w *launcher) addGlobalEnvVars(werk *pb.WerkerTask, builder builderinterface.Builder) {
	acct, repo, _ := stringbuilder.GetAcctRepo(werk.FullName)
	// we don't care if there is an error retrieving this, if it fails it'll return an empty value
	// and that's what we want!
	lastSuccessfulHash, _ := w.Store.GetLastSuccessfulBuildHash(acct, repo, werk.Branch)
	paddedEnvs := []string{
		fmt.Sprintf("GIT_HASH=%s", werk.CheckoutHash),
		fmt.Sprintf("BUILD_ID=%d", werk.Id),
		fmt.Sprintf("GIT_HASH_SHORT=%s", werk.CheckoutHash[:7]),
		fmt.Sprintf("GIT_BRANCH=%s", werk.Branch),
		fmt.Sprintf("WORKSPACE=%s", w.Basher.CloneDir(werk.CheckoutHash)),
		fmt.Sprintf("GIT_PREVIOUS_SUCCESSFUL_COMMIT=%s", lastSuccessfulHash),
	}
	paddedEnvs = append(paddedEnvs, werk.BuildConf.Env...)
	builder.SetGlobalEnv(paddedEnvs)
}

func (w *launcher) listenForDockerUuid(dockerChan chan string, checkoutHash string) error {
	dockerUuid := <-dockerChan

	if err := buildmonitor.RegisterBuild(w.RemoteConf.GetConsul(), w.Uuid.String(), checkoutHash, dockerUuid); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't register build")
		return err
	}

	return nil
}

//storeStageToDb is a helper function for storing stages to db - this runs on completion of every stage
func storeStageToDb(storage storage.BuildStage, buildId int64, stageResult *pb.Result, start time.Time, dur float64) error {
	err := storage.AddStageDetail(&models.StageResult{
		BuildId:       buildId,
		Stage:         stageResult.Stage,
		Status:        int(stageResult.Status),
		Error:         stageResult.Error,
		Messages:      stageResult.Messages,
		StartTime:     start,
		StageDuration: dur,
	})

	if err != nil {
		return err
	}

	return nil
}

func handleFailure(result *pb.Result, store storage.OcelotStorage, stageName string, duration time.Duration, id int64) {
	errStr := fmt.Sprintf("%s stage failed", stageName)
	if len(result.Error) > 0 {
		errStr = errStr + result.Error
	}
	ocelog.Log().Error(errStr)
	if err := store.UpdateSum(true, duration.Seconds(), id); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't update summary in database")
	}
}
