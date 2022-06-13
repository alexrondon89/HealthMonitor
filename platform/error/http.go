package error

import "net/http"

type Error interface {
	Message() string
	Code() int
}

type serviceUnavailableError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

type serviceInternalError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

type customError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

type badRequestError struct {
	Msg      string `json:"message,omitempty"`
	CodeHttp int    `json:"code,omitempty"`
}

func ServiceUnavailable(message string) *serviceUnavailableError {
	return &serviceUnavailableError{
		Msg:      message,
		CodeHttp: http.StatusServiceUnavailable,
	}
}

func ServiceInternal(message string) *serviceInternalError {
	return &serviceInternalError{
		Msg:      message,
		CodeHttp: http.StatusInternalServerError,
	}
}

func Custom(message string, code int) *customError {
	return &customError{
		Msg:      message,
		CodeHttp: code,
	}
}

func BadRequest(message string) *badRequestError {
	return &badRequestError{
		Msg:      message,
		CodeHttp: http.StatusBadRequest,
	}
}

func (su *serviceUnavailableError) Message() string {
	return su.Msg
}

func (si *serviceInternalError) Message() string {
	return si.Msg
}

func (br *badRequestError) Message() string {
	return br.Msg
}

func (ce *customError) Message() string {
	return ce.Msg
}

func (su *serviceUnavailableError) Code() int {
	return su.CodeHttp
}

func (si *serviceInternalError) Code() int {
	return si.CodeHttp
}

func (ce *customError) Code() int {
	return ce.CodeHttp
}

func (br *badRequestError) Code() int {
	return br.CodeHttp
}
