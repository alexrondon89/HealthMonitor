package local

import (
	localDb "HealthMonitor/platform/db/local"
	"HealthMonitor/server/service/repository"
	"fmt"
)

type local struct {
	monitors          localDb.Monitors
	criticalResources localDb.CriticalResources
}

func New() *local {
	return &local{
		monitors:          localDb.LocalMonitor(),
		criticalResources: localDb.LocalCriticalResources(),
	}
}

func (l *local) SaveMonitor(req repository.Input) error {
	value := map[string]string{req.Name: req.Handle}
	l.monitors[req.Type] = value
	fmt.Println("monitores....", l.monitors)
	return nil
}

func (l *local) SaveCriticalResource(name string) error {
	l.criticalResources = append(l.criticalResources, name)
	fmt.Println("criticalResources....", l.criticalResources)
	return nil
}
