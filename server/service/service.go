package service

import "HealthMonitor/platform/error"

type HealthMonitorRegister interface {
	Register(*Request) (*Response, error.Error)
}

type HealthMonitorCheck interface {
	Check() (*Response, error.Error)
}

type Request struct {
	Type     *string `json:"type" validate:"required"`
	Name     *string `json:"name" validate:"required"`
	Handle   *string `json:"handle" validate:"required"`
	Critical *bool   `json:"critical" validate:"required"`
}

type Response struct {
	ClientResponses []*ClientResponse `json:"clients,omitempty"`
	Failed          []string          `json:"failed,omitempty"`
	Message         string            `json:"message,omitempty"`
}

type ClientResponse struct {
	ResourceName string `json:"resourceName,omitempty"`
	Code         int    `json:"code,omitempty"`
	Failed       bool   `json:"failed,omitempty"`
	Message      string `json:"message,omitempty"`
}
