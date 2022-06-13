package errors

import "net/http"

type Error interface {
	Message() string
	Code() int
}

type serviceUnavailableError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

func ServiceUnavailableError(message string) *serviceUnavailableError {
	return &serviceUnavailableError{
		Msg:      message,
		CodeHttp: http.StatusServiceUnavailable,
	}
}

func (su *serviceUnavailableError) Message() string {
	return su.Msg
}

func (su *serviceUnavailableError) Code() int {
	return su.CodeHttp
}

type serviceInternalError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

func ServiceInternalError(message string) *serviceInternalError {
	return &serviceInternalError{
		Msg:      message,
		CodeHttp: http.StatusInternalServerError,
	}
}

func (si *serviceInternalError) Message() string {
	return si.Msg
}

func (si *serviceInternalError) Code() int {
	return si.CodeHttp
}

type customError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

func CustomError(message string, code int) *customError {
	return &customError{
		Msg:      message,
		CodeHttp: code,
	}
}

func (ce *customError) Message() string {
	return ce.Msg
}

func (ce *customError) Code() int {
	return ce.CodeHttp
}

type badRequestError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

func BadRequestError(message string) *badRequestError {
	return &badRequestError{
		Msg:      message,
		CodeHttp: http.StatusBadRequest,
	}
}

func (br *badRequestError) Message() string {
	return br.Msg
}

func (br *badRequestError) Code() int {
	return br.CodeHttp
}
