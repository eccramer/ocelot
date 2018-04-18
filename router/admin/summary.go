package admin

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/level11consulting/go-til/log"

	"bitbucket.org/level11consulting/ocelot/models/pb"
)

func (g *guideOcelotServer) LastFewSummaries(ctx context.Context, repoAct *pb.RepoAccount) (*pb.Summaries, error) {
	log.Log().Debug("getting last few summaries")
	var summaries = &pb.Summaries{}
	var modelz []*pb.BuildSummary
	var err error
	if repoAct.Repo == "" && repoAct.Account == "" {
		modelz, err = g.Storage.RetrieveLastFewSumsAll(repoAct.Limit)
	} else {
		modelz, err = g.Storage.RetrieveLastFewSums(repoAct.Repo, repoAct.Account, repoAct.Limit)
	}
	if err != nil {
		return nil, handleStorageError(err)
	}
	log.Log().Debug("successfully retrieved last few summaries")
	if len(modelz) == 0 {
		return nil, status.Error(codes.NotFound, "no entries found")
	}
	summaries.Sums = modelz
	return summaries, nil

}
