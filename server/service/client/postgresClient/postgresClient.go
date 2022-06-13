package postgresClient

import (
	"github.com/sirupsen/logrus"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/client"
)

type postgresClient struct {
}

func New() *postgresClient {
	return &postgresClient{}
}

func (p *postgresClient) Ping(resourceName string) (*client.Response, error.Error) {
	logrus.Info(`checking resource `, resourceName, `....`)
	return nil, error.ServiceUnavailable("timeout")
}
