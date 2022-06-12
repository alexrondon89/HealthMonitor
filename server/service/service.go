package service

type HealthMonitor interface {
	Register(*Request) error //monitor
	Check() (*Response, error)
}

type Request struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Handle   string `json:"handle"`
	Critical bool   `json:"critical"`
}

type Response struct {
	ClientResponses []*ClientResponses
	Failed          []string
	Status          int
}

type ClientResponses struct {
	ResourceName string
	Code         int
	Failed       bool
	Message      string
}
