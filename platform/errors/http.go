package errors

type CustomError struct {
	Message string
	Code    int
}

func (ce *CustomError) Error() string {
	return ce.Message
}
