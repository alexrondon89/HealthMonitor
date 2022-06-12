package errors

import "net/http"

type Error interface {
	GetMessage() string
	GetCode() int
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

func (su *serviceUnavailableError) GetMessage() string {
	return su.Msg
}

func (su *serviceUnavailableError) GetCode() int {
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

func (si *serviceInternalError) GetMessage() string {
	return si.Msg
}

func (si *serviceInternalError) GetCode() int {
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

func (ce *customError) GetMessage() string {
	return ce.Msg
}

func (ce *customError) GetCode() int {
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

func (br *badRequestError) GetMessage() string {
	return br.Msg
}

func (br *badRequestError) GetCode() int {
	return br.CodeHttp
}
