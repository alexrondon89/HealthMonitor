package postgresClient

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
)

type postgresClient struct {
}

func New() *postgresClient {
	return &postgresClient{}
}

func (p *postgresClient) Ping(resourceName string) (*client.Response, *errors.CustomError) {
	return nil, &errors.CustomError{
		Message: "timeout",
		Code:    503,
	}
}
