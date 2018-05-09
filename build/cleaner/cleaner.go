package cleaner

import (
	"context"

	"github.com/shankj3/ocelot/models"
)

//this interface handles build cleanup
type Cleaner interface {

	//Cleanup performs build cleanup functions. If an optional logout channel is passed, logs will be sent over the channel
	Cleanup(ctx context.Context, id string, logout chan []byte) error
}

//returns a new cleaner interface
func GetNewCleaner(facts *models.WerkerFacts) Cleaner {
	switch facts.WerkerType {
	case models.Docker:
		return &DockerCleaner{}
	case models.Kubernetes:
		return &K8Cleaner{}
	case models.SSH:
		return &SSHCleaner{SSHFacts: facts.Ssh}
	case models.Exec:
		return NewExecCleaner()
	default:
		return &DockerCleaner{}
	}
	return nil
}
