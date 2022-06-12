package repository

import "HealthMonitor/platform/errors"

type Repository interface {
	SaveCriticalResource(*Resource) errors.Error
	SaveMonitor(*Resource) errors.Error
	GetMonitors() (*Monitors, errors.Error)
}

type Resource struct {
	Type   string
	Name   string
	Handle string
}

type Monitors struct {
	Item []Resource
}
