package services

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"fmt"
	"net/http"
	"time"
)

func (s *services) HealthCheck() *transport.Response {
	var (
		dataResponse transport.Response
	)

	// Define a starting time (e.g., the beginning of your health check operation).
	startTime := time.Now()

	// Simulate a health check operation (you can replace this with your own logic).
	// For this example, we'll just sleep for a few seconds.
	time.Sleep(5 * time.Second)

	// Calculate the duration of the health check operation.
	duration := time.Since(startTime)

	if duration.Seconds() > 10 {
		dataResponse = transport.ApiResponse("Running App failed", http.StatusOK, constant.StatusSuccess, fmt.Sprintf("error: %v", duration.Seconds()))
	} else {
		dataResponse = transport.ApiResponse("Running well", http.StatusOK, constant.StatusSuccess, nil)
	}

	return &dataResponse
}
