package local

import (
	"HealthMonitor/platform/db/local"
	"HealthMonitor/platform/errors"
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

func (lr *localRepository) SaveMonitor(input *repository.Resource) errors.Error {
	err := lr.db.SaveMonitor(input.Handle, input.Name, input.Type)
	if err != nil {
		return errors.ServiceInternalError(err.Error())
	}

	return nil
}

func (lr *localRepository) SaveCriticalResource(input *repository.Resource) errors.Error {
	err := lr.db.SaveCriticalResources(input.Name)
	if err != nil {
		return errors.ServiceInternalError(err.Error())
	}

	return nil
}

func (lr *localRepository) GetMonitors() (*repository.Monitors, errors.Error) {
	items, err := lr.db.GetMonitors()
	if err != nil {
		return nil, errors.ServiceInternalError(err.Error())
	}

	if len(items) == 0 {
		return nil, errors.CustomError("there are not resources to check", 200)
	}

	var resp []repository.Resource
	for _, item := range items {
		monitor := repository.Resource{
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
