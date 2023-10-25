package services

import (
	"cakeplabs-technical-test/repositories"
	"cakeplabs-technical-test/transport"
)

type services struct {
	repo repositories.Repositories
}

type Services interface {
	HealthCheck() *transport.Response

	GetMenuById(id int) (*transport.Response, error)
	GetAllMenusByKeyword(keywords string) *transport.Response

	CreateOrder(orderRequest transport.OrderRequest) (*transport.Response, error)
	GetOrderById(id int) (*transport.Response, error)
	GetAllOrdersByKeyword(keywords string) *transport.Response

	GetAdditionById(id int) (*transport.Response, error)
	GetAllAdditionsByKeyword(keywords string) *transport.Response

	SeedData() (*transport.Response, error)
}

func NewServices(r repositories.Repositories) Services {
	return &services{repo: r}
}
