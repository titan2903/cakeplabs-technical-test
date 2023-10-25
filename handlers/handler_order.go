package handlers

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *handler) CreateOrder(c echo.Context) error {
	order := new(transport.OrderRequest)
	c.Bind(order)
	_, err := h.service.CreateOrder(*order)

	if err != nil {
		response := transport.ApiResponse("Failed create order", http.StatusBadRequest, constant.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := transport.ApiResponse("Order created", http.StatusCreated, constant.StatusSuccess, nil)

	return c.JSON(http.StatusCreated, response)
}

func (h *handler) GetOrderById(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response := transport.ApiResponse("Invalid ID", http.StatusBadRequest, constant.StatusError, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	result, err := h.service.GetOrderById(idInt)
	if err != nil {
		response := transport.ApiResponse("Order not found", http.StatusNotFound, constant.StatusError, err)
		return c.JSON(http.StatusNotFound, response)
	}

	response := transport.ApiResponse("Order found", http.StatusOK, constant.StatusSuccess, result.Data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllOrdersByKeyword(c echo.Context) error {
	result := h.service.GetAllOrdersByKeyword(c.QueryParam("keywords"))

	response := transport.ApiResponse("Orders found", http.StatusOK, constant.StatusSuccess, result.Data)
	if result.Data == 0 {
		response = transport.ApiResponse("No orders yet", http.StatusOK, constant.StatusSuccess, transport.OrderMapping{})
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, response)
}
