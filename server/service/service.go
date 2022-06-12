package service

import "HealthMonitor/platform/errors"

type HealthMonitorRegister interface {
	Register(*Request) (*Response, errors.Error)
}

type HealthMonitorCheck interface {
	Check() (*Response, errors.Error)
}
type Request struct {
	Type     string `json:"type,required"`
	Name     string `json:"name,required"`
	Handle   string `json:"handle,required"`
	Critical bool   `json:"critical,required"`
}

type Response struct {
	ClientResponses []*ClientResponses `json:"clients,omitempty"`
	Failed          []string           `json:"failed,omitempty"`
	Message         string             `json:"message,omitempty"`
}

type ClientResponses struct {
	ResourceName string `json:"resourceName,omitempty"`
	Code         int    `json:"code,omitempty"`
	Failed       bool   `json:"failed,omitempty"`
	Message      string `json:"message,omitempty"`
}
