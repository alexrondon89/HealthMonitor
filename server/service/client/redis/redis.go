package redis

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/client"
)

type redis struct {
}

func New() *redis {
	return &redis{}
}

func (r *redis) Ping(resourceName string) (*client.Response, error.Error) {
	logrus.Info(`checking resource `, resourceName, `....`)
	return &client.Response{
		Status: "ok",
		Code:   http.StatusOK,
	}, nil
}
