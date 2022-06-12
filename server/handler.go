package server

import (
	"HealthMonitor/platform/errors"
	"HealthMonitor/server/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
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
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		bodyResponse := errors.ServiceInternalError(err.Error())
		buildResponse(rw, bodyResponse, bodyResponse.GetCode())
		return
	}

	if err := validateRequest(&req); err != nil {
		buildResponse(rw, err, err.GetCode())
		return
	}

	resp, err := h.srvRegister.Register(&req)
	if err != nil {
		buildResponse(rw, err, err.GetCode())
		return
	}

	buildResponse(rw, resp, http.StatusCreated)
	return
}

func (h *handler) HealthCheck(rw http.ResponseWriter, r *http.Request) {
	resp, err := h.srvCheck.Check()
	if err != nil {
		buildResponse(rw, err, err.GetCode())
		return
	}

	if len(resp.Failed) > 0 {
		buildResponse(rw, resp, http.StatusPartialContent)
		return
	}

	buildResponse(rw, resp, http.StatusOK)
	return
}

func validateRequest(req *service.Request) errors.Error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return errors.BadRequestError(err.Error())
	}
	return nil
}
func buildResponse(rw http.ResponseWriter, body interface{}, httpStatus int) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	json.NewEncoder(rw).Encode(body)
}
