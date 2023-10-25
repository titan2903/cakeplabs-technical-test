package handlers

import (
	"cakeplabs-technical-test/services"

	"github.com/labstack/echo"
)

type handler struct {
	service services.Services
}

type Handlers interface {
	HealthCheck(c echo.Context) error

	GetMenuById(c echo.Context) error
	GetAllMenusByKeyword(c echo.Context) error

	GetOrderById(c echo.Context) error
	GetAllOrdersByKeyword(c echo.Context) error
	CreateOrder(c echo.Context) error

	GetAdditionById(c echo.Context) error
	GetAllAdditionsByKeyword(c echo.Context) error

	SeedData(c echo.Context) error
}

func NewHandlers(service services.Services) Handlers {
	return &handler{service: service}
}
