package remote

import (
	"errors"

	"github.com/shankj3/ocelot/common/remote/bitbucket"
	"github.com/shankj3/ocelot/models/bitbucket/pb"
	"github.com/shankj3/ocelot/models/pb"
)

// token calls is for one off calls with an already set token
func SetBuildStatus(vcsType pb.SubCredType, token, hash, repo, account string, buildId int64, status protos.BBState) error {
	switch vcsType {
	case pb.SubCredType_BITBUCKET:
		return bitbucket.SetBuildStatus(token, hash, repo, account, buildId, status)
	default:
		return errors.New("not implemented")
	}
}