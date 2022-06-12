package redis

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
)

type redis struct {
}

func New() *redis {
	return &redis{}
}

func (r *redis) Ping(resourceName string) (*client.Response, *errors.CustomError) {
	return &client.Response{
		Status:       "ok",
		Code:         200,
		ResourceName: resourceName,
	}, nil
}
