package status

import (
	"context"
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/shankj3/ocelot/client/commandhelper"
	models "github.com/shankj3/ocelot/models/pb"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"time"
)

const synopsis = "show status of specific acctname/repo, repo or hash"
const help = `
Usage: ocelot status 
	-build-id <build-id> [optional] if specified, this has first priority
	-hash <hash> [optional] if specified, this has second priority
	-acct-repo <acctname/repo> [optional] if specified, this has third priority
	-repo <repo> [optional] returns status of all repos starting with value
`

func New(ui cli.Ui) *cmd {
	// suppress ui here because there's an ordering to status and the error messages that come stock
	// with OcyHelper may be confusing
	c := &cmd{UI: ui, config: commandhelper.Config, OcyHelper: &commandhelper.OcyHelper{SuppressUI: true}}
	c.init()
	return c
}

type cmd struct {
	UI     cli.Ui
	flags  *flag.FlagSet
	config *commandhelper.ClientConfig
	wide   bool
	buildId int64
	*commandhelper.OcyHelper
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

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	//we accept all 3 flags, but prioritize output in the following order: hash, acct-repo, acct
	c.flags.StringVar(&c.OcyHelper.Hash, "hash", "ERROR", "[optional]  <hash> to display build status")
	c.flags.StringVar(&c.OcyHelper.Repo, "repo", "ERROR", "[optional]  <repo> to display build status")
	c.flags.StringVar(&c.OcyHelper.AcctRepo, "acct-repo", "ERROR", "[optional]  <account>/<repo> to display build status")
	c.flags.BoolVar(&c.wide, "wide", false, "[optional] -wide to see full status description even if build passed")
	c.flags.Int64Var(&c.buildId, "build-id", 0, "[optional] <build-id> to display build status")
}

func (c *cmd) writeStatusErr(err error) {
	status, ok := grpcStatus.FromError(err)
	// if we can't parse the status, just return the shitty error.
	if !ok {
		c.UI.Error(err.Error())
	}
	if status.Code() == codes.NotFound {
		var qualifier string
		if c.Hash != "ERROR" {
			qualifier = c.Hash
		} else if c.AcctRepo != "ERROR" {
			qualifier = c.AcctRepo
		} else if c.Repo != "ERROR" {
			qualifier = c.Repo
		} else if c.buildId != 0 {
			qualifier = fmt.Sprintf("build with id %d", c.buildId)
		}
		c.UI.Error(fmt.Sprintf("Status for %s was not found in the database. It may have not been processed yet.", qualifier))
	} else {
		// here we should post to admin
		c.UI.Error("Error retrieving status, message: " + status.Message())
	}
}

func (c *cmd) Run(args []string) int {
	if err2 := c.flags.Parse(args); err2 != nil {
		c.UI.Error(err2.Error())
		return 1
	}
	// if nothing is set, attempt to detect hash
	if c.OcyHelper.AcctRepo == "ERROR" && c.OcyHelper.Repo == "ERROR" && c.OcyHelper.Hash == "ERROR" && c.buildId == 0 {
		if err := c.OcyHelper.DetectHash(c.UI); err != nil {
			commandhelper.Debuggit(c.UI, err.Error())
			c.UI.Error("You must either be in the repository you want to track, one of the following flags must be set: -acct-repo, -repo, -hash. see --help")
			return 1
		}
	}

	ctx := context.Background()
	if err := commandhelper.CheckConnection(c, ctx); err != nil {
		return 1
	}
	// set the query fields based on the flags that have been set on the status
	var query *models.StatusQuery
	switch {
	// respect build-id first
	case c.buildId != 0:
		commandhelper.Debuggit(c.UI, "using build id for status")
		query = &models.StatusQuery{BuildId: c.buildId}
	// respect set hash second
	case c.OcyHelper.Hash != "ERROR" && len(c.OcyHelper.Hash) > 0:
		commandhelper.Debuggit(c.UI, "using hash for status")
		query = &models.StatusQuery{ Hash: c.OcyHelper.Hash }
	// respect acct-repo third
	case c.OcyHelper.AcctRepo != "ERROR":
		commandhelper.Debuggit(c.UI, "using acct/repo for status")
		if err := c.OcyHelper.SplitAndSetAcctRepo(c.UI); err != nil {
			return 1
		}

		query = &models.StatusQuery{
			AcctName: c.OcyHelper.Account,
			RepoName: c.OcyHelper.Repo,
		}
	// lastly, work from the -repo or detected repo flag
	case c.OcyHelper.Repo != "ERROR":
		commandhelper.Debuggit(c.UI, "using repo for status")
		query = &models.StatusQuery{
			PartialRepo: c.OcyHelper.Repo,
		}
	default:
		panic("unsupported check, this should have been caught by the error checks that happened earlier...")
	}
	var statuses *models.Status
	var err error
	statuses, err = c.GetClient().GetStatus(ctx, query)
	if err != nil {
		c.writeStatusErr(err)
		return 1
	}
	// set the colors an booleans that will determine the description of each stage
	// failedValidation is if it never passed the validation stage, so it was never put on the build queue
	failedValidation := statuses.BuildSum.BuildTime.Seconds == 0 && statuses.BuildSum.BuildTime.Nanos == 0 && statuses.BuildSum.QueueTime.Seconds == 0
	// queued is waiting to be picked up by a werker
	queued := statuses.BuildSum.BuildTime.Seconds == 0 && statuses.BuildSum.BuildTime.Nanos == 0 && statuses.BuildSum.QueueTime.Seconds > 0
	// buildStarted has been picked up
	buildStarted := statuses.BuildSum.BuildTime.Seconds > 0 && statuses.IsInConsul
	// finished is... kinda obvious
	finished := !statuses.IsInConsul && statuses.BuildSum.BuildTime.Seconds > 0
	commandhelper.Debuggit(c.UI, fmt.Sprintf("finished is %v, buildStarted is %v, queued is %v, buildTime is %v", finished, buildStarted, queued, time.Unix(statuses.BuildSum.BuildTime.Seconds, 0)))
	//statuses.BuildSum.QueueTime time.Unix(0,0)
	stageStatus, color, statuss := commandhelper.PrintStatusStages(commandhelper.GetStatus(queued, buildStarted, finished, failedValidation), statuses, c.wide, c.config.Theme)
	buildStatus := commandhelper.PrintStatusOverview(color, statuses.BuildSum.Account, statuses.BuildSum.Repo, statuses.BuildSum.Hash, statuss, c.config.Theme)
	c.UI.Output(buildStatus + stageStatus)
	return 0
}
