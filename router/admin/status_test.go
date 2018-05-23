package admin

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/shankj3/go-til/consul"
	"github.com/shankj3/go-til/test"
	"github.com/shankj3/ocelot/common/credentials"
	"github.com/shankj3/ocelot/models"
	"github.com/shankj3/ocelot/models/pb"
	"github.com/shankj3/ocelot/storage"
)

func TestGuideOcelotServer_GetStatus_retrieveLogic(t *testing.T) {
	ctx := context.Background()
	store := &sumStatusStorage{returnErr:true}
	gos := guideOcelotServer{Storage: store}
	// send query by build id, should retrieve summary by build id and flip the byBuildId flag
	gos.GetStatus(ctx, &pb.StatusQuery{BuildId:123})
	if store.byBuildId == false {
		t.Error("should have retrieved summary by build id")
	}
	store.reset()
	gos.GetStatus(ctx, &pb.StatusQuery{Hash: "123"})
	if store.latestSum == false {
		t.Error("should have retrieved by latest summary by hash")
	}
	store.reset()
	gos.GetStatus(ctx, &pb.StatusQuery{PartialRepo:"ay"})
	if store.lastFewSums == false {
		t.Error("should have retrieved last few sums")
	}
	store.reset()
	gos.GetStatus(ctx, &pb.StatusQuery{AcctName:"1", RepoName:"2"})
	if store.lastFewSums == false {
		t.Error("should have retrieved by last few sums")
	}
	store.reset()
	_, err := gos.GetStatus(ctx, &pb.StatusQuery{AcctName:"1"})
	if err == nil {
		t.Error("shouldreturn an error since only sent status query by account")
	}
}

func TestGuideOcelotServer_GetStatus(t *testing.T) {
	ctx := context.Background()
	store := &sumStatusStorage{returnMulti:true, returnErr: false}
	store.generateTestData()
	gos := guideOcelotServer{Storage: store, RemoteConfig:&credentials.RemoteConfig{}}
	gos.RemoteConfig.SetConsul(&consu{})
	_, err := gos.GetStatus(ctx, &pb.StatusQuery{BuildId:9})
	if err != nil {
		t.Error("even though return multiple summaries is set, this query is by build id so should not be relevant. error is: " + err.Error())
	}
	store.reset()
	_, err = gos.GetStatus(ctx, &pb.StatusQuery{AcctName:"account", RepoName:"repo"})
	if err == nil {
		t.Error("returned more than one build summary and querying by acct & repo, should return error")
	}
	if !strings.Contains(err.Error(), "there is no ONE entry that matches the acctname") {
		t.Error(test.StrFormatErrors("err msg", "there is no ONE entry that matches the acctname .. etc etc", err.Error()))
	}
	// change to returning no build summaries
	store.returnMulti = false
	store.returnNone = true
	store.generateTestData()
	_, err = gos.GetStatus(ctx, &pb.StatusQuery{AcctName:"account", RepoName:"repo"})
	if err == nil {
		t.Error("no summaries were returned, should fail")
	}
	if !strings.Contains(err.Error(), "no entries that match the acctname/repo") {
		t.Error("wrong err msg, err msg is: " + err.Error())
	}
	_, err = gos.GetStatus(ctx, &pb.StatusQuery{PartialRepo:"rep"})
	if err == nil {
		t.Error("no summaries were returned, shoudl fail")
	}
	if !strings.Contains(err.Error(), "there are no repositories starting with ") {
		t.Error("wrong err msg, err msg is: " + err.Error() + "\n and should contain there are no repositories starting with ")
	}

}

// test data/structs //

var summaryGood = models.BuildSummary{
	Hash: "hash",
	Failed: false,
	QueueTime: time.Now(),
	BuildTime: time.Now(),
	Account: "account",
	Repo: "repo",
	Branch: "branch",
	BuildId: 1,
}


var summaryFail = models.BuildSummary{
	Hash: "hash",
	Failed: true,
	QueueTime: time.Now(),
	BuildTime: time.Now(),
	Account: "account",
	Repo: "repo",
	Branch: "branch",
	BuildId: 1,
}


var stagesGood = []models.StageResult{
	{
		BuildId: 1,
		StageResultId: 1,
		Stage: "first",
		Status: int(pb.StageResultVal_PASS),
		Messages: []string{"nice"},
		StartTime: time.Now().Add(-time.Duration(time.Second)),
		StageDuration: 1,
	},
	{
		BuildId: 1,
		StageResultId: 12,
		Stage: "second",
		Status: int(pb.StageResultVal_PASS),
		Messages: []string{"nice2"},
		StartTime: time.Now().Add(-time.Duration(time.Second)),
		StageDuration: 2,
	},
}



var stagesFail = []models.StageResult{
	{
		BuildId: 1,
		StageResultId: 1,
		Stage: "first",
		Status: int(pb.StageResultVal_PASS),
		Messages: []string{"nice"},
		StartTime: time.Now().Add(-time.Duration(time.Second)),
		StageDuration: 1,
	},
	{
		BuildId: 1,
		StageResultId: 12,
		Stage: "second",
		Status: int(pb.StageResultVal_FAIL),
		Messages: []string{"nice2"},
		StartTime: time.Now().Add(-time.Duration(time.Second)),
		StageDuration: 2,
	},
}


type sumStatusStorage struct {
	// if called RetrieveLatestSum
	latestSum bool
	// if called RetrieveLastFewSums
	lastFewSums bool
	// RetrieveSumByBuildId
	byBuildId bool
	returnErr bool
	// if true, return multiple build summaries where appropriate
	returnMulti bool
	// if true, return no build summaries
	returnNone bool
	failed bool

	// generated data
	testSummary models.BuildSummary
	testSummaries []models.BuildSummary
	stages []models.StageResult
	storage.OcelotStorage
}

// generateTestData will loook at the set test flags and generate data appropriately
// if failed is set, then the testSummary field will have a fail status and the stages will also return the last one failed
// if returnMulti is set, then testSummaries will have a length > 1
func (s *sumStatusStorage) generateTestData() {
	if s.failed {
		s.testSummary = summaryFail
		s.stages = stagesFail
	} else {
		s.testSummary = summaryGood
		s.stages = stagesGood
	}
	switch {
	case s.returnNone:
		s.testSummaries = []models.BuildSummary{}
	case s.returnMulti:
		s.testSummaries = []models.BuildSummary{s.testSummary, s.testSummary}
	default:
		s.testSummaries = []models.BuildSummary{s.testSummary}
	}
}


func (s *sumStatusStorage) reset() {
	s.lastFewSums = false
	s.latestSum = false
	s.byBuildId = false
}

func (s *sumStatusStorage) RetrieveSumByBuildId(buildId int64) (models.BuildSummary, error) {
	s.byBuildId = true
	if s.returnErr {
		return models.BuildSummary{}, errors.New("womp womp")
	}
	return models.BuildSummary{}, nil
}

func (s *sumStatusStorage) RetrieveLatestSum(gitHash string) (models.BuildSummary, error) {
	s.latestSum = true
	if s.returnErr {
		return models.BuildSummary{}, errors.New("womp womp")
	}
	return models.BuildSummary{}, nil
}

func (s *sumStatusStorage) RetrieveLastFewSums(repo string, account string, limit int32) ([]models.BuildSummary, error) {
	s.lastFewSums = true
	if s.returnErr {
		return nil, errors.New("womp womp RetrieveLastFewSums")
	}
	return s.testSummaries, nil
}

func (s *sumStatusStorage) RetrieveStageDetail(buildId int64) ([]models.StageResult, error) {
	if s.returnErr {
		return nil, errors.New("womp womp stage detail")
	}
	return s.stages, nil
}

func (s *sumStatusStorage) 	RetrieveAcctRepo(partialRepo string) ([]models.BuildSummary, error) {
	return []models.BuildSummary{{Account:"1", Repo:"2"}}, nil
}

type consu struct {
	returnErr bool
	consul.Consuletty
}

func (c *consu) GetKeyValue(path string) (*api.KVPair, error) {
	if c.returnErr {
		return nil, errors.New("womp womp")
	}
	return &api.KVPair{Key:"dmmy", Value: []byte("dummy")}, nil
}
