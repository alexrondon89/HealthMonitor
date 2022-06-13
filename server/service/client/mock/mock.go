package mock

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/client"
	"github.com/stretchr/testify/mock"
)

type mockClient struct {
	mock.Mock
}

func New() *mockClient {
	return &mockClient{}
}

func (m *mockClient) Ping(resourceName string) (*client.Response, errors.Error) {
	args := m.Called(resourceName)
	if len(m.ExpectedCalls) == 0 {
		return nil, nil
	}
	if args.Get(0) == nil {
		return nil, args.Get(1).(errors.Error)
	}
	return args.Get(0).(*client.Response), nil
}
