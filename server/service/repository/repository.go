package repository

import "HealthMonitor/platform/error"

type Repository interface {
	SaveCriticalResource(*Monitor) error.Error
	SaveMonitor(*Monitor) error.Error
	GetMonitors() (*Monitors, error.Error)
}

type Monitor struct {
	Type   string
	Name   string
	Handle string
}

type Monitors struct {
	Item []Monitor
}
