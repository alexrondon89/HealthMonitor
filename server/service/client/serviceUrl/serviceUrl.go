package serviceUrl

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
	"github.com/sirupsen/logrus"
	"net/http"
)

type serviceUrl struct {
}

func New() *serviceUrl {
	return &serviceUrl{}
}

func (r *serviceUrl) Ping(resourceName string) (*client.Response, errors.Error) {
	logrus.Info("checking resource ", resourceName, "....")
	return &client.Response{
		Status: "ok",
		Code:   http.StatusOK,
	}, nil
}
