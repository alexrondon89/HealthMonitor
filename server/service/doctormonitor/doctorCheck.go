package doctormonitor

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service"
	"HealthMonitor/server/service/client"
	"HealthMonitor/server/service/repository"
	"fmt"
	"sync"
)

type doctorMonitor struct {
	clients map[string]client.Client
	local   repository.Repository
}

func NewChecker(clients map[string]client.Client, local repository.Repository) *doctorMonitor {
	return &doctorMonitor{
		clients: clients,
		local:   local,
	}
}

func (dm *doctorMonitor) Check() (*service.Response, errors.Error) {
	monitors, err := dm.local.GetMonitors()
	if err != nil {
		return nil, err
	}

	n := len(monitors.Item)
	channel := make(chan *service.ClientResponses, n)
	var wg sync.WaitGroup
	wg.Add(n)

	for _, monitor := range monitors.Item {
		go dm.checkClientsHealth(channel, monitor, &wg)
	}

	wg.Wait()
	close(channel)

	return dm.buildServiceResponse(channel)
}

func (dm *doctorMonitor) checkClientsHealth(channel chan<- *service.ClientResponses, monitor repository.Resource, sync *sync.WaitGroup) {
	defer sync.Done()

	if _, ok := dm.clients[monitor.Type]; !ok {
		err := errors.ServiceInternalError(fmt.Sprintf("client not exist for monitor %s", monitor.Type))
		channel <- &service.ClientResponses{
			ResourceName: monitor.Type,
			Failed:       true,
			Code:         err.GetCode(),
			Message:      err.GetMessage(),
		}
		return
	}

	clientResponse, err := dm.clients[monitor.Type].Ping(monitor.Name)
	if err != nil {
		channel <- &service.ClientResponses{
			ResourceName: monitor.Type,
			Failed:       true,
			Code:         err.GetCode(),
			Message:      err.GetMessage(),
		}
		return
	}

	channel <- &service.ClientResponses{
		ResourceName: monitor.Type,
		Code:         clientResponse.Code,
		Message:      clientResponse.Status,
	}
}

func (dm *doctorMonitor) buildServiceResponse(channel <-chan *service.ClientResponses) (*service.Response, errors.Error) {
	var response service.Response
	for cliResp := range channel {
		response.ClientResponses = append(response.ClientResponses, cliResp)
		if cliResp.Failed {
			response.Failed = append(response.Failed, cliResp.ResourceName)
		}
	}

	return &response, nil
}
