package elastic

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
)

type elastic struct {
}

func New() *elastic {
	return &elastic{}
}

func (e *elastic) Ping(resourceName string) (*client.Response, *errors.CustomError) {
	return &client.Response{
		Status:       "ok",
		Code:         200,
		ResourceName: resourceName,
	}, nil
}
