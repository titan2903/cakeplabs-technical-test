package handlers

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *handler) GetAdditionById(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response := transport.ApiResponse("Invalid ID", http.StatusBadRequest, constant.StatusError, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	result, err := h.service.GetAdditionById(idInt)
	if err != nil {
		response := transport.ApiResponse("Addition not found", http.StatusNotFound, constant.StatusError, err)
		return c.JSON(http.StatusNotFound, response)
	}

	response := transport.ApiResponse("Addition found", http.StatusOK, constant.StatusSuccess, result.Data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllAdditionsByKeyword(c echo.Context) error {
	result := h.service.GetAllAdditionsByKeyword(c.QueryParam("keywords"))

	response := transport.ApiResponse("Additions found", http.StatusOK, constant.StatusSuccess, result.Data)
	if result.Data == 0 {
		response = transport.ApiResponse("There is no addition list yet", http.StatusOK, constant.StatusSuccess, transport.MenuMapping{})
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, response)
}
