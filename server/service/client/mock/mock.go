package mock

import (
	"github.com/stretchr/testify/mock"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/client"
)

type mockClient struct {
	mock.Mock
}

func New() *mockClient {
	return &mockClient{}
}

func (m *mockClient) Ping(resourceName string) (*client.Response, error.Error) {
	args := m.Called(resourceName)
	if len(m.ExpectedCalls) == 0 {
		return nil, nil
	}
	if args.Get(0) == nil {
		return nil, args.Get(1).(error.Error)
	}
	return args.Get(0).(*client.Response), nil
}
