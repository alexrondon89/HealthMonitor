package client

import "HealthMonitor/platform/error"

type Client interface {
	Ping(string) (*Response, error.Error)
}

type Response struct {
	Status string
	Code   int
}
