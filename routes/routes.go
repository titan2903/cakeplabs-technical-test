package routes

import (
	"cakeplabs-technical-test/database/config"
	"cakeplabs-technical-test/handlers"
	"cakeplabs-technical-test/repositories"
	"cakeplabs-technical-test/services"

	m "cakeplabs-technical-test/middlewares"

	"github.com/labstack/echo"
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)

	repository := repositories.NewRepositories(config.GetQuery())
	service := services.NewServices(repository)
	handler := handlers.NewHandlers(service)

	e.Validator = m.NewCustomValidator()

	//! Main Routes
	e.GET("/v1/healthcheck", handler.HealthCheck)

	e.GET("/v1/menus/:id", handler.GetMenuById)
	e.GET("/v1/menus", handler.GetAllMenusByKeyword)

	e.GET("/v1/orders/:id", handler.GetOrderById)
	e.GET("/v1/orders", handler.GetAllOrdersByKeyword)
	e.POST("/v1/orders", handler.CreateOrder)

	e.GET("/v1/additions/:id", handler.GetAdditionById)
	e.GET("/v1/additions", handler.GetAllAdditionsByKeyword)

	e.POST("/v1/seeder", handler.SeedData)

	return e
}
