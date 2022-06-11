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
	Status int
	Failed []string
}
