package doctormonitor_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/client"
	mockcli "HealthMonitor/server/service/client/mock"
	"HealthMonitor/server/service/doctormonitor"
	"HealthMonitor/server/service/repository"
	"HealthMonitor/server/service/repository/mock"
)

func TestCheckService(t *testing.T) {
	monitorType := "serviceUrl"
	name := "graphql"
	handler := "www.facebook.com"

	t.Run("monitor must be checked successfully", func(t *testing.T) {
		repositoryMock := mock.New()
		clientMock := mockcli.New()
		cliMockMap := map[string]client.Client{"serviceUrl": clientMock}
		checker := doctormonitor.NewChecker(cliMockMap, repositoryMock)
		monitor := repository.Monitor{Type: monitorType, Name: name, Handle: handler}
		monitors := &repository.Monitors{Item: []repository.Monitor{monitor}}
		cliResp := &client.Response{Status: "ok", Code: http.StatusOK}
		repositoryMock.On("GetMonitors").Return(monitors, nil)
		clientMock.On("Ping", monitor.Name).Return(cliResp, nil)

		response, err := checker.Check()

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.ClientResponses)
		assert.Equal(t, 1, len(response.ClientResponses))
		assert.Equal(t, monitor.Type, response.ClientResponses[0].ResourceName)
		assert.Equal(t, cliResp.Code, response.ClientResponses[0].Code)
		assert.Equal(t, cliResp.Status, response.ClientResponses[0].Message)
		assert.Nil(t, response.Failed)
	})

	t.Run("method GetMonitors must fail getting monitors", func(t *testing.T) {
		repositoryMock := mock.New()
		checker := doctormonitor.NewChecker(nil, repositoryMock)
		msgExpected := "mock error"
		codeExpected := 400
		repositoryMock.On("GetMonitors").Return(nil, error.Custom(msgExpected, codeExpected))

		response, err := checker.Check()

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, msgExpected, err.Message())
		assert.Equal(t, codeExpected, err.Code())

	})

	t.Run("method checkClientsHealth must fail due to not exist client for monitor type", func(t *testing.T) {
		repositoryMock := mock.New()
		clientMock := mockcli.New()
		cliMockMap := map[string]client.Client{"serviceUrl": clientMock}
		checker := doctormonitor.NewChecker(cliMockMap, repositoryMock)
		monitor := repository.Monitor{Type: monitorType + "-fake", Name: name, Handle: handler}
		monitors := &repository.Monitors{Item: []repository.Monitor{monitor}}
		msgExpected := "client not exist for monitor serviceUrl-fake"
		codeExpected := 500
		repositoryMock.On("GetMonitors").Return(monitors, nil)

		response, err := checker.Check()

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, len(response.ClientResponses))
		assert.Equal(t, monitor.Type, response.ClientResponses[0].ResourceName)
		assert.Equal(t, codeExpected, response.ClientResponses[0].Code)
		assert.Equal(t, msgExpected, response.ClientResponses[0].Message)
		assert.NotNil(t, response.Failed)
		assert.Equal(t, 1, len(response.Failed))
	})

	t.Run("method Ping in client throw error for a monitor type", func(t *testing.T) {
		repositoryMock := mock.New()
		clientMock := mockcli.New()
		cliMockMap := map[string]client.Client{"serviceUrl": clientMock}
		checker := doctormonitor.NewChecker(cliMockMap, repositoryMock)
		monitor := repository.Monitor{Type: monitorType, Name: name, Handle: handler}
		monitors := &repository.Monitors{Item: []repository.Monitor{monitor}}
		cliResp := error.ServiceUnavailable("mock error from client")

		repositoryMock.On("GetMonitors").Return(monitors, nil)
		clientMock.On("Ping", monitor.Name).Return(nil, cliResp)

		response, err := checker.Check()

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, len(response.ClientResponses))
		assert.Equal(t, monitor.Type, response.ClientResponses[0].ResourceName)
		assert.Equal(t, cliResp.Code(), response.ClientResponses[0].Code)
		assert.Equal(t, cliResp.Message(), response.ClientResponses[0].Message)
		assert.NotNil(t, response.Failed)
		assert.Equal(t, 1, len(response.Failed))
	})
}
