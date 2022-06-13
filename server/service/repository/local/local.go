package local

import (
	"net/http"

	"HealthMonitor/platform/db/local"
	"HealthMonitor/platform/error"
	"HealthMonitor/server/service/repository"
)

type localRepository struct {
	db *local.DB
}

func New(db *local.DB) *localRepository {
	return &localRepository{
		db: db,
	}
}

func (lr *localRepository) SaveMonitor(input *repository.Monitor) error.Error {
	if err := lr.db.SaveMonitor(input.Handle, input.Name, input.Type); err != nil {
		return error.ServiceInternal(err.Error())
	}

	return nil
}

func (lr *localRepository) SaveCriticalResource(input *repository.Monitor) error.Error {
	if err := lr.db.SaveCriticalResources(input.Name); err != nil {
		return error.ServiceInternal(err.Error())
	}

	return nil
}

func (lr *localRepository) GetMonitors() (*repository.Monitors, error.Error) {
	items, err := lr.db.GetMonitors()
	if err != nil {
		return nil, error.ServiceInternal(err.Error())
	}

	if len(items) == 0 {
		return nil, error.Custom(`there are not resources to check`, http.StatusOK)
	}

	var resp []repository.Monitor
	for _, item := range items {
		monitor := repository.Monitor{
			Type:   item.Type,
			Name:   item.Name,
			Handle: item.Handle,
		}
		resp = append(resp, monitor)
	}

	return &repository.Monitors{
		Item: resp,
	}, nil
}
