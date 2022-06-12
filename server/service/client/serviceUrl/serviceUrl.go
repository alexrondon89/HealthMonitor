package serviceUrl

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
)

type serviceUrl struct {
}

func New() *serviceUrl {
	return &serviceUrl{}
}

func (r *serviceUrl) Ping(resourceName string) (*client.Response, *errors.CustomError) {
	return &client.Response{
		Status:       "ok",
		Code:         200,
		ResourceName: resourceName,
	}, nil
}
