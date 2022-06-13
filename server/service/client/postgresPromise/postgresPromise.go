package postgresPromise

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/client"
)

type postgresPromise struct {
}

func New() *postgresPromise {
	return &postgresPromise{}
}

func (p *postgresPromise) Ping(resourceName string) (*client.Response, error.Error) {
	logrus.Info(`checking resource `, resourceName, `....`)
	return &client.Response{
		Status: "ok",
		Code:   http.StatusOK,
	}, nil
}
