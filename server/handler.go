package server

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service"
	"encoding/json"
	"net/http"
)

type handler struct {
	service service.HealthMonitor
}

type Response struct {
	Message string
}

func New(srv service.HealthMonitor) *handler {
	return &handler{
		service: srv,
	}
}

func (h *handler) ResourceRegister(rw http.ResponseWriter, r *http.Request) {
	var req service.Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		bodyResponse := errors.CustomError{Message: err.Error(), Code: 500}
		buildResponse(rw, bodyResponse, bodyResponse.Code)
		return
	}
	defer r.Body.Close()

	err := h.service.Register(&req)
	if err != nil {
		buildResponse(rw, err, 500)
		return
	}

	resp := Response{Message: "monitor added..."}
	buildResponse(rw, resp, 200)
	return
}

func (h *handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//_ = h.service.Check()
}

func buildResponse(rw http.ResponseWriter, body interface{}, httpStatus int) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	json.NewEncoder(rw).Encode(body)
}
