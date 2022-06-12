package elastic

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
	"github.com/sirupsen/logrus"
	"net/http"
)

type elastic struct {
}

func New() *elastic {
	return &elastic{}
}

func (e *elastic) Ping(resourceName string) (*client.Response, errors.Error) {
	logrus.Info("checking resource ", resourceName, "....")
	return &client.Response{
		Status: "ok",
		Code:   http.StatusOK,
	}, nil
}
