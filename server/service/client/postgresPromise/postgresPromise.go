package postgresPromise

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
)

type postgresPromise struct {
}

func New() *postgresPromise {
	return &postgresPromise{}
}

func (p *postgresPromise) Ping(resourceName string) (*client.Response, *errors.CustomError) {
	return &client.Response{
		Status:       "ok",
		Code:         200,
		ResourceName: resourceName,
	}, nil
}
