package errors

import "net/http"

type Error interface {
	Message() string
	Code() int
}

type serviceUnavailableError struct {
	message string
	code    int
}

func ServiceUnavailableError(message string) *serviceUnavailableError {
	return &serviceUnavailableError{
		message: message,
		code:    http.StatusServiceUnavailable,
	}
}

func (su *serviceUnavailableError) Message() string {
	return su.message
}

func (su *serviceUnavailableError) Code() int {
	return su.code
}

type serviceInternalError struct {
	message string
	code    int
}

func ServiceInternalError(message string) *serviceInternalError {
	return &serviceInternalError{
		message: message,
		code:    http.StatusInternalServerError,
	}
}

func (su *serviceInternalError) Message() string {
	return su.message
}

func (su *serviceInternalError) Code() int {
	return su.code
}

type customError struct {
	message string
	code    int
}

func CustomError(message string, code int) *customError {
	return &customError{
		message: message,
		code:    code,
	}
}

func (su *customError) Message() string {
	return su.message
}

func (su *customError) Code() int {
	return su.code
}
