package poll

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shankj3/go-til/log"
	"github.com/level11consulting/orbitalci/models/pb"
	"github.com/level11consulting/orbitalci/storage"
	"github.com/level11consulting/orbitalci/server/config"
	stringbuilder "github.com/level11consulting/orbitalci/build/helpers/stringbuilder/accountrepo"
	"github.com/shankj3/go-til/nsqpb"
)

type PollSchedule interface {
	PollRepo(context.Context, *pb.PollRequest) (*empty.Empty, error)
	DeletePollRepo(context.Context, *pb.PollRequest) (*empty.Empty, error)
	ListPolledRepos(context.Context, *empty.Empty) (*pb.Polls, error)
}

type PollScheduleAPI struct {
	PollSchedule
	RemoteConfig   config.CVRemoteConfig
	Storage        storage.OcelotStorage
	Producer       nsqpb.Producer
}

func (g *PollScheduleAPI) PollRepo(ctx context.Context, poll *pb.PollRequest) (*empty.Empty, error) {
	if poll.Account == "" || poll.Repo == "" || poll.Cron == "" || poll.Branches == "" || poll.Type == pb.SubCredType_NIL_SCT {
		return nil, status.Error(codes.InvalidArgument, "account, repo, cron, branches, and type are required fields")
	}
	log.Log().Info("recieved poll request for ", poll.Account, poll.Repo, poll.Cron)
	empti := &empty.Empty{}
	exists, err := g.Storage.PollExists(poll.Account, poll.Repo)
	if err != nil {
		return empti, status.Error(codes.Unavailable, "unable to retrieve poll table from storage. err: "+err.Error())
	}
	if exists == true {
		log.Log().Info("updating poll in db")
		if err = g.Storage.UpdatePoll(poll.Account, poll.Repo, poll.Cron, poll.Branches); err != nil {
			msg := "unable to update poll in storage"
			log.IncludeErrField(err).Error(msg)
			return empti, status.Error(codes.Unavailable, msg+": "+err.Error())
		}
	} else {
		log.Log().Info("inserting poll in db")
		creddy, err := config.GetVcsCreds(g.Storage, stringbuilder.CreateAcctRepo(poll.Account, poll.Repo), g.RemoteConfig, poll.Type)
		if err != nil {
			var msg string
			if _, ok := err.(*storage.ErrMultipleVCSTypes); ok {
				msg = "multiple vcs types for this account, please include the Type"
			} else {
				msg = "unable to find credentials for account " + poll.Account
			}
			log.IncludeErrField(err).Error(msg)
			return empti, status.Error(codes.InvalidArgument, msg+": "+err.Error())
		}
		if err = g.Storage.InsertPoll(poll.Account, poll.Repo, poll.Cron, poll.Branches, creddy.GetId()); err != nil {
			msg := "unable to insert poll into storage"
			log.IncludeErrField(err).Error(msg)
			return empti, status.Error(codes.Unavailable, msg+": "+err.Error())
		}
	}
	log.Log().WithField("account", poll.Account).WithField("repo", poll.Repo).Info("successfully added/updated poll in storage")
	err = g.Producer.WriteProto(poll, "poll_please")
	if err != nil {
		log.IncludeErrField(err).Error("couldn't write to queue producer at poll_please")
		return empti, status.Error(codes.Unavailable, err.Error())
	}
	return empti, nil
}

func (g *PollScheduleAPI) DeletePollRepo(ctx context.Context, poll *pb.PollRequest) (*empty.Empty, error) {
	if poll.Account == "" || poll.Repo == "" {
		return nil, status.Error(codes.InvalidArgument, "account and repo are required fields")
	}
	log.Log().Info("received delete poll request for ", poll.Account, " ", poll.Repo)
	empti := &empty.Empty{}
	if err := g.Storage.DeletePoll(poll.Account, poll.Repo); err != nil {
		log.IncludeErrField(err).WithField("account", poll.Account).WithField("repo", poll.Repo).Error("couldn't delete poll")
	}
	log.Log().WithField("account", poll.Account).WithField("repo", poll.Repo).Info("successfully deleted poll in storage")
	if err := g.Producer.WriteProto(poll, "no_poll_please"); err != nil {
		log.IncludeErrField(err).Error("couldn't write to queue producer at no_poll_please")

		return empti, status.Error(codes.Unavailable, err.Error())
	}
	return empti, nil
}

// todo: add acct/repo action later
func (g *PollScheduleAPI) ListPolledRepos(context.Context, *empty.Empty) (*pb.Polls, error) {
	polls, err := g.Storage.GetAllPolls()
	if err != nil {
		if _, ok := err.(*storage.ErrNotFound); !ok {
			return nil, status.Error(codes.Unavailable, err.Error())
		}
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &pb.Polls{Polls: polls}, nil
}
