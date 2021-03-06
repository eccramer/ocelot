package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/level11consulting/orbitalci/build/buildeventhandler/push/buildjob"
	"github.com/level11consulting/orbitalci/build/buildeventhandler/push/webhook"
	"github.com/level11consulting/orbitalci/client/buildconfigvalidator"
	"github.com/level11consulting/orbitalci/client/newbuildjob"
	"github.com/level11consulting/orbitalci/models/pb"
	"github.com/level11consulting/orbitalci/build/commiteventhandler"
	"github.com/level11consulting/orbitalci/server/config"
	"github.com/level11consulting/orbitalci/version"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shankj3/go-til/deserialize"
	ocelog "github.com/shankj3/go-til/log"
	ocenet "github.com/shankj3/go-til/net"
	"github.com/shankj3/go-til/nsqpb"
)

// FIXME: consistency: consul's host and port, the var name for configInstance/remoteConfig
func main() {
	//ocelog.InitializeLog("debug")
	defaultName, _ := os.Hostname()

	var consulHost, loglevel, name string
	var consulPort int
	flrg := flag.NewFlagSet("hookhandler", flag.ExitOnError)

	flrg.StringVar(&name, "name", defaultName, "if wish to identify as other than hostname")
	flrg.StringVar(&consulHost, "consul-host", "localhost", "host / ip that consul is running on")
	flrg.StringVar(&loglevel, "log-level", "info", "log level")
	flrg.IntVar(&consulPort, "consul-port", 8500, "port that consul is running on")
	flrg.Parse(os.Args[1:])
	version.MaybePrintVersion(flrg.Args())
	ocelog.InitializeLog(loglevel)
	ocelog.Log().Debug()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
		ocelog.Log().Warning("running on default port :8088")
	}

	parsedConsulURL, parsedErr := url.Parse(fmt.Sprintf("consul://%s:%d", consulHost, consulPort))
	if parsedErr != nil {
		ocelog.IncludeErrField(parsedErr).Fatal("failed parsing consul uri, bailing")
	}

	remoteConfig, err := config.GetInstance(parsedConsulURL, "")
	if err != nil {
		ocelog.Log().Fatal(err)
	}

	//mode := os.Getenv("ENV")
	//if strings.EqualFold(mode, "dev") {
	//	hookHandlerContext = &hh.MockHookHandlerContext{}
	//	hookHandlerContext.SetRemoteConfig(&hh.MockRemoteConfig{})
	//	ocelog.Log().Info("hookhandler running in dev mode")
	//
	//} else {
	store, err := remoteConfig.GetOcelotStorage()
	if err != nil {
		ocelog.IncludeErrField(err).Fatal("couldn't get storage!")
	}
	signaler := &buildjob.Signaler{RC: remoteConfig, Deserializer: deserialize.New(), Producer: nsqpb.GetInitProducer(), OcyValidator: buildconfigvalidator.GetOcelotValidator(), Store: store}
	hookHandlerContext := commiteventhandler.GetContext(signaler, &newbuildjob.PushWerkerTeller{}, &webhook.PullReqWerkerTeller{})
	defer store.Close()

	startServer(hookHandlerContext, port)
}

func startServer(ctx *commiteventhandler.HookHandlerContext, port string) {
	muxi := mux.NewRouter()

	// handleBBevent can take push/pull/ w/e
	muxi.HandleFunc("/"+strings.ToLower(pb.SubCredType_BITBUCKET.String()), ctx.HandleBBEvent).Methods("POST")
	muxi.HandleFunc("/"+strings.ToLower(pb.SubCredType_GITHUB.String()), ctx.HandleGHEvent).Methods("POST")
	muxi.Handle("/metrics", promhttp.Handler())
	n := ocenet.InitNegroni("hookhandler", muxi)
	n.Run(":" + port)
}
