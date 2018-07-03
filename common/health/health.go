/*
health is a package for generating the /health endpoint for all services in ocelot
*/
package health

import (
	"os"
	"path"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	"github.com/shankj3/go-til/consul"
	"github.com/shankj3/go-til/vault"
	"github.com/shankj3/ocelot/models/pb"
	"github.com/shankj3/ocelot/storage"
	"github.com/shankj3/ocelot/version"
)

// Generate will
func Generate(storage storage.OcelotStorage, storageSignificance string, consuletty consul.Consuletty, consulSignificance string, vaulty vault.Vaulty, vaultSignificance string, startTime time.Time) (*pb.Health, error) {
	now := time.Now()
	stampedNow := &timestamp.Timestamp{Seconds:now.Unix()}
	stampedStart := &timestamp.Timestamp{Seconds: startTime.Unix()}
	// get name of service by the executable name
	binary, err := os.Executable()
	if err != nil {
		return nil, errors.WithMessage(err, "could not get path to executable for service name")
	}
	_, serviceName := path.Split(binary)
	var dependencyHealths []*pb.DependencyHealth
	var healthy = true
	// set healthy components for db
	var status string
	if storage != nil {
		if !storage.Healthy() {
			status = "Red"
			healthy = false
		} else {
			status = "Green"
		}
		dependencyHealths = append(dependencyHealths, &pb.DependencyHealth{
			Name: "storage",
			Url: storage.Detail(),
			Status: status,
			Significance: storageSignificance,
			LastUpdated: stampedNow,
		})

	}
	if consuletty != nil {
		// set healthy components for consul
		if !consuletty.IsConnected() {
			status = "Red"
			healthy = false
		} else {
			status = "Green"
		}
		dependencyHealths = append(dependencyHealths, &pb.DependencyHealth{
			Name: "consul",
			Url: consuletty.Detail(),
			Status: status,
			Significance: consulSignificance,
			LastUpdated: stampedNow,
		})
	}

	if vaulty != nil {
		// vault
		if !vaulty.Healthy() {
			healthy = false
			status = "Red"
		} else {
			status = "Green"
		}
		dependencyHealths = append(dependencyHealths, &pb.DependencyHealth{
			Name: "vault",
			Url: vaulty.GetAddress(),
			Status: status,
			Significance: vaultSignificance,
			LastUpdated: stampedNow,
		})
		if !healthy {
			status = "Red"
		} else {
			status = "Green"
		}
	}
	return &pb.Health{
		Status: status,
		Description: serviceName,
		StartTime: stampedStart,
		Version: version.GetHumanVersion(),
		Dependencies: dependencyHealths,
	}, nil
}