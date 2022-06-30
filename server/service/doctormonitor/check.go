package doctormonitor

import (
	"HealthMonitor/platform/error"
	"HealthMonitor/server/service"
	"HealthMonitor/server/service/client"
	"HealthMonitor/server/service/repository"
	"sync"
)

type doctorCheck struct {
	clients map[string]client.Client
	local   repository.Repository
}

func NewChecker(clients map[string]client.Client, local repository.Repository) *doctorCheck {
	return &doctorCheck{
		clients: clients,
		local:   local,
	}
}

func (dm *doctorCheck) Check() (*service.Response, error.Error) {
	monitors, err := dm.local.GetMonitors()
	if err != nil {
		return nil, err
	}

	n := len(monitors.Item)
	wg := &sync.WaitGroup{}
	wg.Add(n)
	channel := make(chan *service.ClientResponse, 1)

	go func() {
		for _, monitor := range monitors.Item {
			go dm.checkClientsHealth(channel, monitor, wg)
		}
		wg.Wait()
		close(channel)
	}()

	return dm.buildResponse(channel)
}

func (dm *doctorCheck) checkClientsHealth(channel chan<- *service.ClientResponse, monitor repository.Monitor, sync *sync.WaitGroup) {
	defer sync.Done()

	if _, ok := dm.clients[monitor.Type]; !ok {
		return
	}

	clientResponse, err := dm.clients[monitor.Type].Ping(monitor.Name)
	if err != nil {
		channel <- &service.ClientResponse{
			ResourceName: monitor.Type,
			Failed:       true,
			Code:         err.Code(),
			Message:      err.Message(),
		}
		return
	}

	channel <- &service.ClientResponse{
		ResourceName: monitor.Type,
		Code:         clientResponse.Code,
		Message:      clientResponse.Status,
	}
}

func (dm *doctorCheck) buildResponse(channel <-chan *service.ClientResponse) (*service.Response, error.Error) {
	response := service.Response{}
	for cliResp := range channel {
		response.ClientResponses = append(response.ClientResponses, cliResp)
		if cliResp.Failed {
			response.Failed = append(response.Failed, cliResp.ResourceName)
		}
	}

	return &response, nil
}
