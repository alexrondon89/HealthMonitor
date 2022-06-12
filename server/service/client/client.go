package client

import "HealthMonitor/platform/errors"

type Client interface {
	Ping(string) (*Response, *errors.CustomError)
}

type Response struct {
	ResourceName string
	Status       string
	Code         int
}
