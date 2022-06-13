package postgresPool

import (
	"github.com/sirupsen/logrus"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/client"
)

type postgresPool struct {
}

func New() *postgresPool {
	return &postgresPool{}
}

func (p *postgresPool) Ping(resourceName string) (*client.Response, error.Error) {
	logrus.Info(`checking resource `, resourceName, `....`)
	return nil, error.ServiceUnavailable("timeout")
}
