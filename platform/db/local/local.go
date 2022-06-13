package local

import (
	"github.com/sirupsen/logrus"
)

type DB struct {
	monitors          []Monitor
	criticalResources []string
}

type Monitor struct {
	Type   string
	Name   string
	Handle string
}

func New() *DB {
	return &DB{
		monitors:          []Monitor{},
		criticalResources: []string{},
	}
}

func (ldb *DB) SaveMonitor(handle string, name string, typ string) error {
	monitor := Monitor{Type: typ, Name: name, Handle: handle}
	ldb.monitors = append(ldb.monitors, monitor)
	logrus.Info(`monitors....`, ldb.monitors)
	return nil
}

func (ldb *DB) SaveCriticalResources(value string) error {
	ldb.criticalResources = append(ldb.criticalResources, value)
	logrus.Info(`criticalResources....`, ldb.monitors)
	return nil
}

func (ldb *DB) GetMonitors() ([]Monitor, error) {
	return ldb.monitors, nil
}
