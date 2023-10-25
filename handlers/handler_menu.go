package handlers

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *handler) GetMenuById(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response := transport.ApiResponse("Invalid ID", http.StatusBadRequest, constant.StatusError, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	result, err := h.service.GetMenuById(idInt)
	if err != nil {
		response := transport.ApiResponse("Menu not found", http.StatusNotFound, constant.StatusError, err)
		return c.JSON(http.StatusNotFound, response)
	}

	response := transport.ApiResponse("Menu found", http.StatusOK, constant.StatusSuccess, result.Data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllMenusByKeyword(c echo.Context) error {
	result := h.service.GetAllMenusByKeyword(c.QueryParam("keywords"))

	response := transport.ApiResponse("Menus found", http.StatusOK, constant.StatusSuccess, result.Data)
	if result.Data == 0 {
		response = transport.ApiResponse("There is no menu list yet", http.StatusOK, constant.StatusSuccess, transport.MenuMapping{})
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, response)
}
