package local

import (
	"fmt"
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
	ldb.monitors = append(ldb.monitors, Monitor{Type: typ, Name: name, Handle: handle})
	fmt.Println("monitores....", ldb.monitors)
	return nil
}

func (ldb *DB) SaveCriticalResources(value string) error {
	ldb.criticalResources = append(ldb.criticalResources, value)
	fmt.Println("criticalResources....", ldb.monitors)

	return nil
}

func (ldb *DB) GetMonitors() ([]Monitor, error) {
	return ldb.monitors, nil
}
