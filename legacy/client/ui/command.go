package client

import (
	"github.com/level11consulting/orbitalci/client/ui/build"
	"github.com/level11consulting/orbitalci/client/ui/creds/apple"
	"github.com/level11consulting/orbitalci/client/ui/creds/apple/applelist"
	"github.com/level11consulting/orbitalci/client/ui/creds/delete"
	"github.com/level11consulting/orbitalci/client/ui/creds/env"
	"github.com/level11consulting/orbitalci/client/ui/creds/env/add"
	"github.com/level11consulting/orbitalci/client/ui/creds/env/list"
	"github.com/level11consulting/orbitalci/client/ui/creds/helmrepo/add"
	"github.com/level11consulting/orbitalci/client/ui/creds/helmrepo/list"
	"github.com/level11consulting/orbitalci/client/ui/creds/k8s"
	"github.com/level11consulting/orbitalci/client/ui/creds/notify"
	"github.com/level11consulting/orbitalci/client/ui/creds/ssh"
	"github.com/level11consulting/orbitalci/client/ui/init"
	"github.com/level11consulting/orbitalci/models/pb"
	"github.com/mitchellh/cli"

	"github.com/level11consulting/orbitalci/client/ui/creds"
	"github.com/level11consulting/orbitalci/client/ui/creds/notify/notifyadd"
	"github.com/level11consulting/orbitalci/client/ui/creds/notify/notifylist"

	//"github.com/level11consulting/orbitalci/client/ui/creds/credsadd"
	//"github.com/level11consulting/orbitalci/client/ui/creds/credslist"
	"github.com/level11consulting/orbitalci/client/ui/creds/apple/appleadd"
	"github.com/level11consulting/orbitalci/client/ui/creds/k8s/add"
	"github.com/level11consulting/orbitalci/client/ui/creds/k8s/list"
	"github.com/level11consulting/orbitalci/client/ui/creds/repocreds"
	"github.com/level11consulting/orbitalci/client/ui/creds/repocreds/add"
	"github.com/level11consulting/orbitalci/client/ui/creds/repocreds/list"
	"github.com/level11consulting/orbitalci/client/ui/creds/ssh/sshadd"
	"github.com/level11consulting/orbitalci/client/ui/creds/ssh/sshlist"
	"github.com/level11consulting/orbitalci/client/ui/creds/vcscreds"
	"github.com/level11consulting/orbitalci/client/ui/creds/vcscreds/add"
	"github.com/level11consulting/orbitalci/client/ui/creds/vcscreds/list"
	"github.com/level11consulting/orbitalci/client/ui/kill"
	"github.com/level11consulting/orbitalci/client/ui/logs"
	"github.com/level11consulting/orbitalci/client/ui/poll/add"
	"github.com/level11consulting/orbitalci/client/ui/poll/delete"
	"github.com/level11consulting/orbitalci/client/ui/poll/list"
	"github.com/level11consulting/orbitalci/client/ui/repos"
	"github.com/level11consulting/orbitalci/client/ui/repos/list"
	"github.com/level11consulting/orbitalci/client/ui/status"
	"github.com/level11consulting/orbitalci/client/ui/summary"
	"github.com/level11consulting/orbitalci/client/ui/validate"
	"github.com/level11consulting/orbitalci/client/ui/version"
	"github.com/level11consulting/orbitalci/client/ui/watch"
	ocyVersion "github.com/level11consulting/orbitalci/version"

	"os"
)

var Commands map[string]cli.CommandFactory

func init() {
	base := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr, Reader: os.Stdin}
	ui := &cli.ColoredUi{Ui: base, OutputColor: cli.UiColorNone, InfoColor: cli.UiColorNone, ErrorColor: cli.UiColorRed, WarnColor: cli.UiColorYellow}
	verHuman := ocyVersion.GetHumanVersion()
	Commands = map[string]cli.CommandFactory{
		"creds": func() (cli.Command, error) { return creds.New(), nil },
		// todo: fix these funct  ions then add them back in
		//"creds add":                func() (cli.Command, error) { return credsadd.New(ui), nil },
		//"creds list":           func() (cli.Command, error) { return credslist.New(ui), nil },
		"creds vcs":              func() (cli.Command, error) { return vcscreds.New(), nil },
		"creds vcs list":         func() (cli.Command, error) { return buildcredslist.New(ui), nil },
		"creds vcs add":          func() (cli.Command, error) { return buildcredsadd.New(ui), nil },
		"creds ssh":              func() (cli.Command, error) { return ssh.New(), nil },
		"creds ssh list":         func() (cli.Command, error) { return sshlist.New(ui), nil },
		"creds ssh add":          func() (cli.Command, error) { return sshadd.New(ui), nil },
		"creds ssh delete":       func() (cli.Command, error) { return delete.New(ui, pb.CredType_SSH), nil },
		"creds repo":             func() (cli.Command, error) { return repocreds.New(), nil },
		"creds repo add":         func() (cli.Command, error) { return repocredsadd.New(ui), nil },
		"creds repo list":        func() (cli.Command, error) { return repocredslist.New(ui), nil },
		"creds repo delete":      func() (cli.Command, error) { return delete.New(ui, pb.CredType_REPO), nil },
		"creds k8s":              func() (cli.Command, error) { return k8s.New(), nil },
		"creds k8s add":          func() (cli.Command, error) { return kubeadd.New(ui), nil },
		"creds k8s list":         func() (cli.Command, error) { return kubelist.New(ui), nil },
		"creds k8s delete":       func() (cli.Command, error) { return delete.New(ui, pb.CredType_K8S), nil },
		"creds apple":            func() (cli.Command, error) { return apple.New(), nil },
		"creds apple add":        func() (cli.Command, error) { return appleadd.New(ui), nil },
		"creds apple list":       func() (cli.Command, error) { return applelist.New(ui), nil },
		"creds notify":           func() (cli.Command, error) { return notify.New(), nil },
		"creds notify add":       func() (cli.Command, error) { return notifyadd.New(ui), nil },
		"creds notify list":      func() (cli.Command, error) { return notifylist.New(ui), nil },
		"creds notify delete":    func() (cli.Command, error) { return delete.New(ui, pb.CredType_NOTIFIER), nil },
		"creds env":              func() (cli.Command, error) { return env.New(), nil },
		"creds env add":          func() (cli.Command, error) { return envadd.New(ui), nil },
		"creds env list":         func() (cli.Command, error) { return envlist.New(ui), nil },
		"creds env delete":       func() (cli.Command, error) { return delete.New(ui, pb.CredType_GENERIC), nil },
		"creds helmrepo add":     func() (cli.Command, error) { return helmrepoadd.New(ui), nil },
		"creds helmrepo list":    func() (cli.Command, error) { return helmrepolist.New(ui), nil },
		"creds helmrepo  delete": func() (cli.Command, error) { return delete.New(ui, pb.CredType_GENERIC), nil },
		"init":                   func() (cli.Command, error) { return ocyinit.New(ui), nil },
		"logs":                   func() (cli.Command, error) { return logs.New(ui), nil },
		"summary":                func() (cli.Command, error) { return summary.New(ui), nil },
		"validate":               func() (cli.Command, error) { return validate.New(ui), nil },
		"status":                 func() (cli.Command, error) { return status.New(ui), nil },
		"watch":                  func() (cli.Command, error) { return watch.New(ui), nil },
		"build":                  func() (cli.Command, error) { return build.New(ui), nil },
		"poll":                   func() (cli.Command, error) { return polladd.New(ui), nil },
		"poll delete":            func() (cli.Command, error) { return polldelete.New(ui), nil },
		"poll list":              func() (cli.Command, error) { return polllist.New(ui), nil },
		"kill":                   func() (cli.Command, error) { return kill.New(ui), nil },
		"version":                func() (cli.Command, error) { return version.New(ui, verHuman), nil },
		"repos":                  func() (cli.Command, error) { return repos.New(), nil },
		"repos list":             func() (cli.Command, error) { return reposlist.New(ui), nil },
	}
}
