package credslist

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/cli"
	"github.com/shankj3/ocelot/client/commandhelper"
	//"github.com/shankj3/ocelot/client/creds/repocreds/list"
	//"github.com/shankj3/ocelot/client/creds/vcscreds/list"
	models "github.com/shankj3/ocelot/models/pb"
)

func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui, config: commandhelper.Config}
	c.init()
	return c
}

type cmd struct {
	UI            cli.Ui
	flags         *flag.FlagSet
	accountFilter string
	config        *commandhelper.ClientConfig
	account       string
}

func (c *cmd) GetClient() models.GuideOcelotClient {
	return c.config.Client
}

func (c *cmd) GetUI() cli.Ui {
	return c.UI
}

func (c *cmd) GetConfig() *commandhelper.ClientConfig {
	return c.config
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.StringVar(&c.account, "account", "",
		"Account to filter credentials on")
}

func writeCred(ui cli.Ui, account string, cred models.OcyCredder) {
	if account == "" || account == cred.GetAcctName() {
		ui.Info(cred.ClientString())
	}
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}
	ctx := context.Background()
	var protoReq empty.Empty
	if err := commandhelper.CheckConnection(c, ctx); err != nil {
		return 1
	}
	fmt.Println("hi")
	//msg, err := c.config.Client.GetAllCreds(ctx, &protoReq)
	//if err != nil {
	//	c.UI.Error(fmt.Sprint("Could not get list of credentials!\n Error: ", err.Error()))
	//	return 1
	//}

	vcs, _ := c.config.Client.GetVCSCreds(ctx, &protoReq)
	repo, _ := c.config.Client.GetRepoCreds(ctx, &protoReq)
	notify, _ := c.config.Client.GetNotifyCreds(ctx, &protoReq)
	env, _ := c.config.Client.GetGenericCreds(ctx, &protoReq)
	ssh, _ := c.config.Client.GetSSHCreds(ctx, &protoReq)
	k8s, _ := c.config.Client.GetK8SCreds(ctx, &protoReq)

	c.UI.Info("---- All Credentials ----\n")
	if vcs != nil {
		c.UI.Info("---- VCS ----")
		for _, cred := range vcs.Vcs {
			writeCred(c.UI, c.account, cred)
		}
	}
	if repo != nil {
		c.UI.Info("---- REPO ----")
		for _, cred := range repo.Repo {
			writeCred(c.UI, c.account, cred)
		}
	}
	if notify != nil {
		c.UI.Info("---- NOTIFY ----")
		for _, cred := range notify.Creds {
			writeCred(c.UI, c.account, cred)
		}
	}
	if env != nil {
		c.UI.Info("---- ENV ----")
		for _, cred := range env.Creds {
			writeCred(c.UI, c.account, cred)
		}
	}
	if ssh != nil {
		c.UI.Info("--- SSH ---")
		for _, cred := range ssh.Keys {
			writeCred(c.UI, c.account, cred)
		}
	}
	if k8s != nil {
		c.UI.Info("--- K8S ---")
		for _, cred := range k8s.K8SCreds {
			writeCred(c.UI, c.account, cred)
		}
	}
	return 0
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

const synopsis = "list all credentials added to ocelot"
const help = `
Usage: ocelot creds list

  Will list all credentials that have been added to ocelot. 
`
