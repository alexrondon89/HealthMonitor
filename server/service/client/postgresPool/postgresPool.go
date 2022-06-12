package postgresPool

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
	"github.com/sirupsen/logrus"
)

type postgresPool struct {
}

func New() *postgresPool {
	return &postgresPool{}
}

func (p *postgresPool) Ping(resourceName string) (*client.Response, errors.Error) {
	logrus.Info("checking resource ", resourceName, "....")
	return nil, errors.ServiceUnavailableError("timeout")
}
