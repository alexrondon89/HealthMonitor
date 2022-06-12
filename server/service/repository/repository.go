package repository

type Repository interface {
	SaveCriticalResource(*Resource) error
	SaveMonitor(*Resource) error
	GetMonitors() (*Monitors, error)
}

type Resource struct {
	Type   string
	Name   string
	Handle string
}

type Monitors struct {
	Item []Resource
}
