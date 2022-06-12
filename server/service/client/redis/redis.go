package redis

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
	"github.com/sirupsen/logrus"
	"net/http"
)

type redis struct {
}

func New() *redis {
	return &redis{}
}

func (r *redis) Ping(resourceName string) (*client.Response, errors.Error) {
	logrus.Info("checking resource ", resourceName, "....")
	return &client.Response{
		Status: "ok",
		Code:   http.StatusOK,
	}, nil
}
