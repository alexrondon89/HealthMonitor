package client

import "HealthMonitor/platform/errors"

type Client interface {
	Ping(string) (*Response, errors.Error)
}

type Response struct {
	Status string
	Code   int
}
