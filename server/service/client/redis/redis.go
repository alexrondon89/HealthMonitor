package redis

type redis struct {
}

func New() *redis {
	return &redis{}
}

func (e *redis) Ping() {

}
