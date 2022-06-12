package doctormonitor

import (
	"HealthMonitor/server/service"
	"HealthMonitor/server/service/client"
	"HealthMonitor/server/service/repository"
	"sync"
)

type doctorMonitor struct {
	clients map[string]client.Client
	local   repository.Repository
}

func New(clients map[string]client.Client, local repository.Repository) *doctorMonitor {
	return &doctorMonitor{
		clients: clients,
		local:   local,
	}
}

func (dm *doctorMonitor) Register(resource *service.Request) error {
	res := &repository.Resource{
		Type:   resource.Type,
		Name:   resource.Name,
		Handle: resource.Handle,
	}
	err := dm.local.SaveMonitor(res)
	if err != nil {
		return err
	}

	if resource.Critical {
		err = dm.local.SaveCriticalResource(res)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dm *doctorMonitor) Check() (*service.Response, error) {
	monitors, err := dm.local.GetMonitors()
	if err != nil {
		return nil, err
	}

	n := len(monitors.Item)
	channel := make(chan *service.ClientResponses, n)
	var wg sync.WaitGroup
	wg.Add(n)

	for _, monitor := range monitors.Item {
		go dm.catchClientResponse(channel, monitor, &wg)
	}

	wg.Wait()
	close(channel)

	return dm.buildServiceResponse(channel)
}

func (dm *doctorMonitor) catchClientResponse(channel chan<- *service.ClientResponses, monitor repository.Resource, sync *sync.WaitGroup) {
	defer sync.Done()
	clientResponse, err := dm.clients[monitor.Type].Ping(monitor.Name)
	if err != nil {
		channel <- &service.ClientResponses{
			ResourceName: monitor.Name,
			Code:         err.Code,
			Failed:       true,
			Message:      err.Message,
		}
	}

	channel <- &service.ClientResponses{
		ResourceName: clientResponse.ResourceName,
		Code:         clientResponse.Code,
		Message:      clientResponse.Status,
	}
}

func (dm *doctorMonitor) buildServiceResponse(channel <-chan *service.ClientResponses) (*service.Response, error) {
	var response service.Response
	for cliResp := range channel {
		response.ClientResponses = append(response.ClientResponses, cliResp)
		if cliResp.Failed {
			response.Failed = append(response.Failed, cliResp.ResourceName)
		}
	}

	if len(response.Failed) > 0 {
		response.Status = 206
	}
	return &response, nil
}
