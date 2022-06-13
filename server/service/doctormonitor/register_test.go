package doctormonitor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service"
	"HealthMonitor/server/service/doctormonitor"
	"HealthMonitor/server/service/repository"
	"HealthMonitor/server/service/repository/mock"
)

func TestRegisterService(t *testing.T) {
	monitorType := "serviceUrl"
	name := "graphql"
	handler := "www.facebook.com"
	critical := true

	t.Run("monitor must be added successfully", func(t *testing.T) {
		repositoryMock := mock.New()
		register := doctormonitor.NewRegistrator(repositoryMock)
		reqSrv := &service.Request{Type: &monitorType, Name: &name, Handle: &handler, Critical: &critical}
		reqRepo := &repository.Monitor{Type: monitorType, Name: name, Handle: handler}
		msgExpected := "monitor serviceUrl with name graphql registered successfully"

		repositoryMock.On("SaveMonitor", reqRepo).Return(nil)
		repositoryMock.On("SaveCriticalResource", reqRepo).Return(nil)

		response, err := register.Register(reqSrv)

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Nil(t, response.ClientResponses)
		assert.Nil(t, response.Failed)
		assert.Equal(t, msgExpected, response.Message)
	})

	t.Run("method SaveMonitor must fail and return error", func(t *testing.T) {
		repositoryMock := mock.New()
		register := doctormonitor.NewRegistrator(repositoryMock)
		reqSrv := &service.Request{Type: &monitorType, Name: &name, Handle: &handler, Critical: &critical}
		reqRepo := &repository.Monitor{Type: monitorType, Name: name, Handle: handler}
		msgExpected := "mock error"
		codeExpected := 400

		repositoryMock.On("SaveMonitor", reqRepo).Return(error.Custom(msgExpected, codeExpected))

		response, err := register.Register(reqSrv)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, msgExpected, err.Message())
		assert.Equal(t, codeExpected, err.Code())

	})

	t.Run("method SaveCriticalResource must fail and return error", func(t *testing.T) {
		repositoryMock := mock.New()
		register := doctormonitor.NewRegistrator(repositoryMock)
		reqSrv := &service.Request{Type: &monitorType, Name: &name, Handle: &handler, Critical: &critical}
		reqRepo := &repository.Monitor{Type: monitorType, Name: name, Handle: handler}
		msgExpected := "mock error"
		codeExpected := 400

		repositoryMock.On("SaveMonitor", reqRepo).Return(nil)
		repositoryMock.On("SaveCriticalResource", reqRepo).Return(error.Custom(msgExpected, codeExpected))

		response, err := register.Register(reqSrv)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, msgExpected, err.Message())
		assert.Equal(t, codeExpected, err.Code())

	})
}
