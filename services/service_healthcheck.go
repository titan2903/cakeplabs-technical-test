package services

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"
)

func (s *services) HealthCheck() *transport.Response {
	dataResponse := transport.ApiResponse("Running well", http.StatusOK, constant.StatusSuccess, nil)
	return &dataResponse
}
