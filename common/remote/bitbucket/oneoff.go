package bitbucket

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	pbb "github.com/shankj3/ocelot/models/bitbucket/pb"
)

func SetBuildStatus(token, hash, repo, account string, buildId int64, status pbb.BBState) error {
		marshaler := jsonpb.Marshaler{}
		urlend := fmt.Sprintf("%s/%s/commit/%s/statuses/build", account, repo, hash)
		url := fmt.Sprintf(DefaultRepoBaseURL, urlend)
		buildStatus := &pbb.Status{
			State: 		  status.String(),
			Key:   		  fmt.Sprintf("OCELOT-BUILD-%d", buildId),
			Url:   		  fmt.Sprintf("https://admin-ocelot.metaverse.l11.com/builds/%s/%s/%s/%d", repo, account, hash, buildId),
			Description:  fmt.Sprintf("ocelot build of %s/%s at commit %s", account, repo, hash),
			Name: 		  fmt.Sprintf("OCELOT-BUILD-%s-%s-%d", repo, account, buildId),
		}
		postData, err := marshaler.MarshalToString(buildStatus)
		if err != nil {
			return err
		}
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, strings.NewReader(postData))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		_, err = client.Do(req)
		return err
	}