package doctormonitor

import (
	"HealthMonitor/server/service"
	"HealthMonitor/server/service/client"
	"HealthMonitor/server/service/repository"
)

type doctorMonitor struct {
	elastic         client.Client
	postgresPromise client.Client
	redis           client.Client
	serviceUrl      client.Client
	local           repository.Repository
}

func New(elastic, postgresPromise, redis, serviceUrl client.Client, local repository.Repository) *doctorMonitor {
	return &doctorMonitor{
		elastic:         elastic,
		postgresPromise: postgresPromise,
		redis:           redis,
		serviceUrl:      serviceUrl,
		local:           local,
	}
}

func (dm *doctorMonitor) Register(resource *service.Request) error {
	res := repository.Input{
		Type:     resource.Type,
		Name:     resource.Name,
		Handle:   resource.Handle,
		Critical: resource.Critical,
	}
	err := dm.local.SaveMonitor(res)
	if err != nil {
		return err
	}

	if res.Critical {
		err = dm.local.SaveCriticalResource(res.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dm *doctorMonitor) Check() (*service.Response, error) {
	return nil, nil
}
