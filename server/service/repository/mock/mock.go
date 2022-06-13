package mock

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service/repository"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func New() *mockRepository {
	return &mockRepository{}
}

func (m *mockRepository) SaveCriticalResource(resource *repository.Resource) errors.Error {
	args := m.Called(resource)
	if len(m.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errors.Error)
}

func (m *mockRepository) SaveMonitor(resource *repository.Resource) errors.Error {
	args := m.Called(resource)
	if len(m.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errors.Error)
}

func (m *mockRepository) GetMonitors() (*repository.Monitors, errors.Error) {
	args := m.Called()
	if len(m.ExpectedCalls) == 0 {
		return nil, nil
	}
	if args.Get(0) == nil {
		return nil, args.Get(1).(errors.Error)
	}
	return args.Get(0).(*repository.Monitors), nil
}
