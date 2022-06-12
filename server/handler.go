package server

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service"
	"encoding/json"
	"net/http"
)

type handler struct {
	srvRegister service.HealthMonitorRegister
	srvCheck    service.HealthMonitorCheck
}

func New(srvRegister service.HealthMonitorRegister, srvCheck service.HealthMonitorCheck) *handler {
	return &handler{
		srvRegister: srvRegister,
		srvCheck:    srvCheck,
	}
}

func (h *handler) ResourceRegister(rw http.ResponseWriter, r *http.Request) {
	var req service.Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		bodyResponse := errors.ServiceInternalError(err.Error())
		buildResponse(rw, bodyResponse, bodyResponse.Code())
		return
	}
	defer r.Body.Close()

	resp, err := h.srvRegister.Register(&req)
	if err != nil {
		buildResponse(rw, err, err.Code())
		return
	}

	buildResponse(rw, resp, http.StatusCreated)
	return
}

func (h *handler) HealthCheck(rw http.ResponseWriter, r *http.Request) {
	resp, err := h.srvCheck.Check()
	if err != nil {
		buildResponse(rw, err, err.Code())
		return
	}

	if len(resp.Failed) > 0 {
		buildResponse(rw, resp, http.StatusPartialContent)
		return
	}

	buildResponse(rw, resp, http.StatusOK)
	return
}

func buildResponse(rw http.ResponseWriter, body interface{}, httpStatus int) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	json.NewEncoder(rw).Encode(body)
}
