package postgresPromise

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
	"github.com/sirupsen/logrus"
	"net/http"
)

type postgresPromise struct {
}

func New() *postgresPromise {
	return &postgresPromise{}
}

func (p *postgresPromise) Ping(resourceName string) (*client.Response, errors.Error) {
	logrus.Info("checking resource ", resourceName, "....")
	return &client.Response{
		Status: "ok",
		Code:   http.StatusOK,
	}, nil
}
