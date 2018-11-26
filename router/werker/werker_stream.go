package werker

import (
	"net/http"
	"path"
	"runtime"

	"github.com/gorilla/websocket"
	ocelog "github.com/shankj3/go-til/log"

)

var (
	upgrader = websocket.Upgrader{}
)

func addHandlers(muxi *http.ServeMux, werkData *WerkerContext) {
	//serve up zip files that spawned containers need
	muxi.HandleFunc("/do_things.tar", func(w http.ResponseWriter, r *http.Request) {
		if werkData.Dev {
			ocelog.Log().Info("DEV MODE, SERVING WERKER FILES LOCALLY")
			_, filename, _, ok := runtime.Caller(0)
			if !ok {
				panic("no caller???? ")
			}
			http.ServeFile(w, r, path.Dir(filename)+"/werker_files.tar")
		} else {
			// todo: move to the base64 encode then echo to file method, this is frustrating
			ocelog.Log().Debug("serving up zip files from s3")
			/// todo: change this back!!!!
			http.Redirect(w, r, "https://s3-us-west-2.amazonaws.com/ocelotty/werker_files_dev.tar", 301)
		}
	})

	// todo: THESE ARE ALL LINUX-SPECIFIC!
	muxi.HandleFunc("/kubectl", func(w http.ResponseWriter, r *http.Request) {
		ocelog.Log().Debug("serving up kubectl binary from googleapis")
		http.Redirect(w, r, "https://storage.googleapis.com/kubernetes-release/release/v1.9.6/bin/linux/amd64/kubectl", 301)
	})
	muxi.HandleFunc("/helm.tar.gz", func(w http.ResponseWriter, r *http.Request) {
		ocelog.Log().Debug("serving up helm binary from googleapis")
		http.Redirect(w, r, "https://storage.googleapis.com/kubernetes-helm/helm-v2.10.0-linux-amd64.tar.gz", 301)
	})
	muxi.HandleFunc("/mc", func(w http.ResponseWriter, r *http.Request) {
		ocelog.Log().Debug("serving up mc binary")
		http.Redirect(w, r, "https://dl.minio.io/client/mc/release/linux-amd64/mc", 301)
	})
}