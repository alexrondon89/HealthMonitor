package postgresClient

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
	"github.com/sirupsen/logrus"
)

type postgresClient struct {
}

func New() *postgresClient {
	return &postgresClient{}
}

func (p *postgresClient) Ping(resourceName string) (*client.Response, errors.Error) {
	logrus.Info("checking resource ", resourceName, "....")
	return nil, errors.ServiceUnavailableError("timeout")
}
