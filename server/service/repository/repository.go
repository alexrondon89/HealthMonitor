package repository

type Repository interface {
	SaveCriticalResource(string) error
	SaveMonitor(Input) error
}

type Input struct {
	Type     string
	Name     string
	Handle   string
	Critical bool
}
