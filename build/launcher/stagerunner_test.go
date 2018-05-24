package launcher

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/shankj3/ocelot/build"
	"github.com/shankj3/ocelot/build/basher"
	"github.com/shankj3/ocelot/build/valet"
	"github.com/shankj3/ocelot/common/credentials"
	"github.com/shankj3/ocelot/models"
	"github.com/shankj3/ocelot/storage"


	"github.com/shankj3/go-til/test"
	"github.com/shankj3/ocelot/models/pb"
)

func Test_handleTriggers(t *testing.T) {
	var triggerData = []struct {
		branch      string
		shouldSkip  bool
		store       *dummyBuildStage
		shouldError bool
	}{
		{"boogaloo", true, &dummyBuildStage{details: []*models.StageResult{}}, false},
		{"alks;djf", true, &dummyBuildStage{details: []*models.StageResult{}, fail: true}, true},
		{"vibranium", false, &dummyBuildStage{details: []*models.StageResult{}}, false},
	}
	triggers := &pb.Triggers{Branches: []string{"apple", "banana", "quartz", "vibranium",}}
	stage := &pb.Stage{Env: []string{}, Script: []string{"echo suuuup yooo"}, Name: "testing_triggers", Trigger: triggers}

	for ind, wd := range triggerData {
		t.Run(fmt.Sprintf("%d-trigger", ind), func(t *testing.T) {
			shouldSkip, err := handleTriggers(wd.branch, 12, wd.store, stage)
			if err != nil && !wd.shouldError {
				t.Fatal(err)
			}
			if wd.shouldError && err == nil {
				t.Error("handleTriggers should have errored, didn't")
			}
			if wd.shouldSkip != shouldSkip {
				t.Logf("branch %s | shouldSkip %v | shouldError %v", wd.branch, wd.shouldSkip, wd.shouldError)
				t.Error(test.GenericStrFormatErrors("should skip", wd.shouldSkip, shouldSkip))
			}
		})
	}
}

func Test_handleTriggers__badRegex(t *testing.T) {
	badTrigger := &pb.Triggers{Branches:[]string{`[\d{3}-\d{3}-\d{4}]`}}
	stage := &pb.Stage{Env: []string{}, Script: []string{"echo suuuup yooo"}, Name: "testing_triggers", Trigger: badTrigger}
	_, err := handleTriggers("goobranch", 12, &dummyBuildStage{details:[]*models.StageResult{}}, stage)
	if err == nil {
		t.Error("should error")
	}
	t.Log(err)
}

func Test_handleTriggers__badStore(t *testing.T) {
	badTrigger := &pb.Triggers{Branches:[]string{`[\d{3}-\d{3}-\d{4}]`}}
	stage := &pb.Stage{Env: []string{}, Script: []string{"echo suuuup yooo"}, Name: "testing_triggers", Trigger: badTrigger}
	_, err := handleTriggers("goobranch", 12, &fakeStore{returnErr:true}, stage)
	if err == nil {
		t.Error("should error")
	}
	if err.Error() != "i was told to fail" {
		t.Error("error message should be from storage, it is instead: " + err.Error())
	}
}

func Test_handleTriggers__nilTrigger(t *testing.T) {
	stage := &pb.Stage{Env: []string{}, Script: []string{"echo suuuup yooo"}, Name: "testing_triggers", Trigger: nil}
	skip, err := handleTriggers("goobranch", 12, &dummyBuildStage{details:[]*models.StageResult{}}, stage)
	if skip || err != nil {
		t.Error("handleTriggers, when called with an empty trigger block, should return a skip of false and a nil error")
	}
}


func Test_handleTriggers__TriggerNoBranches(t *testing.T) {
	stage := &pb.Stage{Env: []string{}, Script: []string{"echo suuuup yooo"}, Name: "testing_triggers", Trigger: &pb.Triggers{Branches:[]string{}}}
	skip, err := handleTriggers("goobranch", 12, &dummyBuildStage{details:[]*models.StageResult{}}, stage)
	if skip || err != nil {
		t.Error("handleTriggers, when called with a trigger block with no branches, should return a skip of false and a nil error")
	}
}

func Test_runStages(t *testing.T) {
	builder := &fakeBuilder{OcyBash: &basher.Basher{}}
	launchr := &launcher{BuildValet: &testValet{}, Store: &fakeStore{returnErr:false}, infochan: make(chan []byte)}
	task := &pb.WerkerTask{
		Id: 1,
		CheckoutHash: "hash",
		Branch: "branch",
		BuildConf: &pb.BuildConfig{
			Stages: []*pb.Stage{
				{
					Name:"test",
					Env: []string{"ENV1=VAL1", "ENV2=VAL2"},
					Script: []string{"echo ayyyyyyy"},
				},
				{
					Name:"test2",
					Env: []string{"ENV3=VAL3", "ENV3=VAL3"},
					Script: []string{"echo traaaaaaayyyyyy"},
				},

			},
		},
	}
	ctx := context.Background()
	fail, _, err := launchr.runStages(ctx, task, builder)
	if err != nil {
		t.Error(err)
	}
	if fail {
		t.Error("booooo, should not have failed")
	}
	if diff := deep.Equal(builder.executedStrings, []string{"echo ayyyyyyy", "echo traaaaaaayyyyyy"}); diff != nil {
		t.Error(diff)
	}

	// set trigger of second stage so taht it doesn't run
	task.BuildConf.Stages[1].Trigger = &pb.Triggers{Branches: []string{"dontrunme"}}
	// reset executedStrings
	builder.executedStrings = []string{}
	fail, _, err = launchr.runStages(ctx, task, builder)
	if err != nil {
		t.Error(err)
	}
	if fail {
		t.Error("booooo, should not have failed")
	}
	// should handle triggers, and not execute second stage
	if diff := deep.Equal(builder.executedStrings, []string{"echo ayyyyyyy"}); diff != nil {
		t.Error(diff)
	}
	//reset
	builder.executedStrings = []string{}
	// make fakeBuilder fail on first stage
	builder.failStage = true
	fail, _, err = launchr.runStages(ctx, task, builder)
	if err != nil {
		t.Error(err)
	}
	if !fail {
		t.Error("should have returned a fail=true as the fake builder failed the first stage")
	}
}



func Test_runStages__badStore(t *testing.T) {
	builder := &fakeBuilder{OcyBash: &basher.Basher{}}
	launchr := &launcher{BuildValet: &testValet{}, Store: &fakeStore{returnErr:true}, infochan: make(chan []byte)}
	task := &pb.WerkerTask{
		Id: 1,
		CheckoutHash: "hash",
		Branch: "branch",
		BuildConf: &pb.BuildConfig{
			Stages: []*pb.Stage{
				{
					Name:"test",
					Env: []string{"ENV1=VAL1", "ENV2=VAL2"},
					Script: []string{"echo ayyyyyyy"},
				},
				{
					Name:"test2",
					Env: []string{"ENV3=VAL3", "ENV3=VAL3"},
					Script: []string{"echo traaaaaaayyyyyy"},
				},

			},
		},
	}
	ctx := context.Background()
	_, _, err := launchr.runStages(ctx, task, builder)
	if err == nil {
		t.Error("should return an error!")
	}
	if err.Error() != "i was told to fail" {
		t.Error("should return the error that the storage implementation threw")
	}
}

/// test structs ///

type fakeBuilder struct {
	build.OcyBash
	envs []string
	executedStrings []string
	failStage bool
}

func (f *fakeBuilder) Init(ctx context.Context, hash string, logout chan []byte) *pb.Result {
	return  &pb.Result{Messages:[]string{}, Status:pb.StageResultVal_PASS}
}

func (f *fakeBuilder) SetGlobalEnv(envs []string) {
	f.envs = envs
}


func (f *fakeBuilder) Setup(ctx context.Context, logout chan []byte, dockerId chan string, werk *pb.WerkerTask, rc credentials.CVRemoteConfig, werkerPort string) (res *pb.Result, uuid string) {
	dockerId <- "fake"
	return &pb.Result{Messages:[]string{}, Status:pb.StageResultVal_PASS}, "fake"
}

func (f *fakeBuilder) Execute(ctx context.Context, actions *pb.Stage, logout chan []byte, commitHash string) *pb.Result {
	f.executedStrings = append(f.executedStrings, strings.Join(actions.Script, " "))
	if f.failStage {
		return &pb.Result{Messages:[]string{"i was told to fail"}, Status:pb.StageResultVal_FAIL}
	}
	return &pb.Result{Messages:[]string{}, Status:pb.StageResultVal_PASS}
}

func (f *fakeBuilder) ExecuteIntegration(ctx context.Context, stage *pb.Stage, stgUtil *build.StageUtil, logout chan []byte) *pb.Result {
	return &pb.Result{Messages:[]string{}, Status:pb.StageResultVal_PASS}
}

func (f *fakeBuilder) GetContainerId() string {
	return "fake"
}

func (f *fakeBuilder) Close() error {
	return nil
}

type testValet struct {
	valet.BuildValet
}

func (tv *testValet)  Reset(newStage string, hash string) error {
	return nil
}



type dummyBuildStage struct {
	details []*models.StageResult
	fail    bool
}

func (dbs *dummyBuildStage) AddStageDetail(stageResult *models.StageResult) error {
	if dbs.fail {
		return errors.New("i am failing as promised")
	}
	dbs.details = append(dbs.details, stageResult)
	return nil
}

func (dbs *dummyBuildStage) RetrieveStageDetail(buildId int64) ([]models.StageResult, error) {
	var srs []models.StageResult
	for _, i := range dbs.details {
		srs = append(srs, *i)
	}
	return srs, nil
}

type fakeStore struct {
	storage.OcelotStorage
	// fail on AddStageDetail if this is true
	returnErr bool
}

func (fs *fakeStore)  AddStageDetail(stageResult *models.StageResult) error {
	if fs.returnErr {
		return errors.New("i was told to fail")
	}
	return nil
}
