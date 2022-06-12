package doctormonitor

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service"
	"HealthMonitor/server/service/repository"
	"fmt"
)

type doctorRegister struct {
	local repository.Repository
}

func NewRegistrator(local repository.Repository) *doctorRegister {
	return &doctorRegister{
		local: local,
	}
}

func (dm *doctorRegister) Register(resource *service.Request) (*service.Response, errors.Error) {
	req := dm.buildRequest(resource)
	err := dm.local.SaveMonitor(req)
	if err != nil {
		return nil, err
	}

	if resource.Critical {
		err = dm.local.SaveCriticalResource(req)
		if err != nil {
			return nil, err
		}
	}

	resp := dm.buildResponse(resource)
	return resp, nil
}

func (dm *doctorRegister) buildRequest(resource *service.Request) *repository.Resource {
	return &repository.Resource{
		Type:   resource.Type,
		Name:   resource.Name,
		Handle: resource.Handle,
	}
}

func (dm *doctorRegister) buildResponse(resource *service.Request) *service.Response {
	return &service.Response{
		Message: fmt.Sprintf("monitor %s with name %s registered successfully", resource.Type, resource.Name),
	}
}
