package admin

import (

	"github.com/shankj3/go-til/deserialize"
	"github.com/shankj3/go-til/nsqpb"

	"github.com/level11consulting/orbitalci/build/helpers/buildscript/validate"
	"github.com/level11consulting/orbitalci/client/buildconfigvalidator"
	"github.com/level11consulting/orbitalci/models/pb"
	"github.com/level11consulting/orbitalci/server/config"
	"github.com/level11consulting/orbitalci/storage"
	"github.com/level11consulting/orbitalci/secret"
	"github.com/level11consulting/orbitalci/secret/anycred"
	"github.com/level11consulting/orbitalci/secret/appledev"
	"github.com/level11consulting/orbitalci/secret/artifactrepo"
	"github.com/level11consulting/orbitalci/secret/generic"
	"github.com/level11consulting/orbitalci/secret/kubernetes"
	"github.com/level11consulting/orbitalci/secret/notifier"
	"github.com/level11consulting/orbitalci/secret/ssh"
	"github.com/level11consulting/orbitalci/secret/vcs"
	"github.com/level11consulting/orbitalci/repo"
	"github.com/level11consulting/orbitalci/repo/poll"
	"github.com/level11consulting/orbitalci/server/grpc/admin/status"
	"github.com/level11consulting/orbitalci/build"
)

//this is our grpc server, it responds to client requests
type OcelotServerAPI struct {
	build.BuildAPI
	repo.RepoInterfaceAPI
	status.StatusInterfaceAPI
	secret.SecretInterfaceAPI
}

func NewOcelotServer(config config.CVRemoteConfig, d *deserialize.Deserializer, adminV *validate.AdminValidator, repoV *validate.RepoValidator, storage storage.OcelotStorage, hhBaseUrl string) pb.GuideOcelotServer {

	anyCredAPI := anycred.AnyCredAPI {
		Storage:        storage,	
		RemoteConfig:   config,
	}

	buildAPI := build.BuildAPI {
		Storage:        storage,	
		RemoteConfig:   config,
		Deserializer:   d,
		Producer:       nsqpb.GetInitProducer(),
		OcyValidator:   buildconfigvalidator.GetOcelotValidator(),
	}

	appleDevSecretAPI := appledev.AppleDevSecretAPI {
		Storage:        storage,	
		RemoteConfig:   config,
	}

	artifactRepoSecretAPI := artifactrepo.ArtifactRepoSecretAPI {
		Storage:        storage,	
		RemoteConfig:   config,
		RepoValidator:  repoV,
	}

	genericSecretAPI := generic.GenericSecretAPI {
		Storage:        storage,	
		RemoteConfig:   config,
	}

	kubernetesSecretAPI := kubernetes.KubernetesSecretAPI {
		Storage:        storage,	
		RemoteConfig:   config,
	}

	notifierSecretAPI := notifier.NotifierSecretAPI {
		Storage:        storage,	
		RemoteConfig:   config,
	}

	pollScheduleAPI := poll.PollScheduleAPI {
		Storage:        storage,	
		RemoteConfig:   config,
		Producer:       nsqpb.GetInitProducer(),
	}

	repoInterfaceAPI := repo.RepoInterfaceAPI {
		PollScheduleAPI:pollScheduleAPI,
		RemoteConfig:   config,
		Storage:        storage,	
		HhBaseUrl:      hhBaseUrl,
	}

	sshSecretAPI := ssh.SshSecretAPI {
		Storage:        storage,	
		RemoteConfig:   config,
	}

	statusInterfaceAPI := status.StatusInterfaceAPI {
		Storage:        storage,	
		RemoteConfig:   config,
	}

	vcsSecretAPI := vcs.VcsSecretAPI {
		Storage:        storage,	
		RemoteConfig:   config,
		AdminValidator: adminV,
	}

	secretInterfaceAPI := secret.SecretInterfaceAPI {
		AnyCredAPI: anyCredAPI,
		AppleDevSecretAPI: appleDevSecretAPI,
		ArtifactRepoSecretAPI: artifactRepoSecretAPI,
		GenericSecretAPI: genericSecretAPI,
		KubernetesSecretAPI: kubernetesSecretAPI,
		NotifierSecretAPI: notifierSecretAPI,
		SshSecretAPI: sshSecretAPI,
		VcsSecretAPI: vcsSecretAPI,
	}

	return &OcelotServerAPI{ 
		buildAPI,
		repoInterfaceAPI,
		statusInterfaceAPI,
		secretInterfaceAPI,
	}
}
