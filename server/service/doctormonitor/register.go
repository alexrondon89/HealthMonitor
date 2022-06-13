package doctormonitor

import (
	"fmt"

	"HealthMonitor/platform/error"
	"HealthMonitor/server/service"
	"HealthMonitor/server/service/repository"
)

type doctorRegister struct {
	local repository.Repository
}

func NewRegistrator(local repository.Repository) *doctorRegister {
	return &doctorRegister{
		local: local,
	}
}

func (dm *doctorRegister) Register(resource *service.Request) (*service.Response, error.Error) {
	req := dm.buildRequest(resource)
	if err := dm.local.SaveMonitor(req); err != nil {
		return nil, err
	}

	if *resource.Critical {
		if err := dm.local.SaveCriticalResource(req); err != nil {
			return nil, err
		}
	}

	return dm.buildResponse(resource)
}

func (dm *doctorRegister) buildRequest(resource *service.Request) *repository.Monitor {
	return &repository.Monitor{
		Type:   *resource.Type,
		Name:   *resource.Name,
		Handle: *resource.Handle,
	}
}

func (dm *doctorRegister) buildResponse(resource *service.Request) (*service.Response, error.Error) {
	return &service.Response{
		Message: fmt.Sprintf(`monitor %s with name %s registered successfully`, *resource.Type, *resource.Name),
	}, nil
}
