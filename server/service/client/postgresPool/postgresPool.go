package postgresPool

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
)

type postgresPool struct {
}

func New() *postgresPool {
	return &postgresPool{}
}

func (p *postgresPool) Ping(resourceName string) (*client.Response, *errors.CustomError) {
	return nil, &errors.CustomError{
		Message: "timeout",
		Code:    503,
	}
}
