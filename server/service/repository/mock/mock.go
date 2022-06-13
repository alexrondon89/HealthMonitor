package mock

import (
	"github.com/stretchr/testify/mock"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/repository"
)

type mockRepository struct {
	mock.Mock
}

func New() *mockRepository {
	return &mockRepository{}
}

func (m *mockRepository) SaveCriticalResource(resource *repository.Monitor) error.Error {
	args := m.Called(resource)
	if len(m.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error.Error)
}

func (m *mockRepository) SaveMonitor(resource *repository.Monitor) error.Error {
	args := m.Called(resource)
	if len(m.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error.Error)
}

func (m *mockRepository) GetMonitors() (*repository.Monitors, error.Error) {
	args := m.Called()
	if len(m.ExpectedCalls) == 0 {
		return nil, nil
	}
	if args.Get(0) == nil {
		return nil, args.Get(1).(error.Error)
	}
	return args.Get(0).(*repository.Monitors), nil
}
