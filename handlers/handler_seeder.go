package handlers

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"

	"github.com/labstack/echo"
)

func (h *handler) SeedData(c echo.Context) error {
	result, err := h.service.SeedData()

	response := transport.ApiResponse("Success seed data menus, toppings and fillings", http.StatusOK, constant.StatusSuccess, result.Data)

	if err != nil {
		response = transport.ApiResponse("Failed seed data menus, toppings and fillings", http.StatusInternalServerError, constant.StatusError, err)
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, response)
}
