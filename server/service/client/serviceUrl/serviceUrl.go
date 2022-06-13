package serviceUrl

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/client"
)

type serviceUrl struct {
}

func New() *serviceUrl {
	return &serviceUrl{}
}

func (r *serviceUrl) Ping(resourceName string) (*client.Response, error.Error) {
	logrus.Info(`checking resource `, resourceName, `....`)
	return &client.Response{
		Status: "ok",
		Code:   http.StatusOK,
	}, nil
}
