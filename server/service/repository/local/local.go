package local

import (
	"HealthMonitor/platform/db/local"
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

func (lr *localRepository) SaveMonitor(input *repository.Resource) error {
	err := lr.db.SaveMonitor(input.Handle, input.Name, input.Type)
	if err != nil {
		return err
	}

	return nil
}

func (lr *localRepository) SaveCriticalResource(input *repository.Resource) error {
	err := lr.db.SaveCriticalResources(input.Name)
	if err != nil {
		return err
	}

	return nil
}

func (lr *localRepository) GetMonitors() (*repository.Monitors, error) {
	items, err := lr.db.GetMonitors()
	if err != nil {
		return nil, err
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
