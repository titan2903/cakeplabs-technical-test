package services

import (
	"cakeplabs-technical-test/entity"
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/pkg/helper"
	"cakeplabs-technical-test/transport"
	"errors"
	"net/http"

	"github.com/labstack/gommon/log"
)

func (s *services) CreateOrder(orderRequest transport.OrderRequest) (*transport.Response, error) {
	var (
		totalPrice   float64
		toppingPrice float64
		fillingPrice float64
		toppingName  string
		fillingName  string
	)

	menu, err := s.repo.GetMenuById(orderRequest.MenuID)
	if err != nil {
		log.Error("GetMenuById order request", err)
		return nil, errors.New("menu not found")
	}

	additionalTopping := s.repo.GetAllAdditionsByKeyword(orderRequest.ToppingName)

	additionalFilling := s.repo.GetAllAdditionsByKeyword(orderRequest.FillingName)

	if len(orderRequest.FillingName) > 0 && helper.DataAddition[orderRequest.FillingName] == orderRequest.FillingName {
		fillingPrice = float64(additionalFilling[0].Price)
		fillingName = additionalFilling[0].Name
	} else {
		return nil, errors.New("filling not found")
	}

	if len(orderRequest.ToppingName) > 0 && helper.DataAddition[orderRequest.ToppingName] == orderRequest.ToppingName {
		toppingPrice = float64(additionalTopping[0].Price)
		toppingName = additionalTopping[0].Name
	} else {
		return nil, errors.New("topping not found")
	}

	//! Calculate order total price
	totalPrice = float64(menu.Price) + toppingPrice + fillingPrice

	dataOrder := entity.Order{
		MenuID:      menu.ID,
		ToppingName: toppingName,
		FillingName: fillingName,
		TotalPrice:  totalPrice,
	}

	err = s.repo.CreateOrder(dataOrder)
	if err != nil {
		log.Error("Failed create order", err)
		return nil, err
	}

	result := transport.ApiResponse("Order created", http.StatusOK, constant.StatusSuccess, dataOrder)
	return &result, nil
}

func (s *services) GetOrderById(id int) (*transport.Response, error) {
	order, err := s.repo.GetOrderById(id)
	if err != nil {
		log.Error("GetOrderById", err)
		return nil, err
	}

	orderMapping := &transport.OrderMapping{
		ID:          order.ID,
		MenuName:    order.Menu.Name,
		ToppingName: order.ToppingName,
		FillingName: order.FillingName,
		TotalPrice:  order.TotalPrice,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}

	result := transport.ApiResponse("Success get order", http.StatusOK, constant.StatusSuccess, orderMapping)
	return &result, nil
}

func (s *services) GetAllOrdersByKeyword(keywords string) *transport.Response {
	orders := s.repo.GetAllOrdersByKeyword(keywords)

	var arrOrders []transport.OrderMapping
	for _, order := range orders {
		orderMapping := transport.OrderMapping{
			ID:          order.ID,
			MenuName:    order.Menu.Name,
			ToppingName: order.ToppingName,
			FillingName: order.FillingName,
			TotalPrice:  order.TotalPrice,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		}

		arrOrders = append(arrOrders, orderMapping)
	}

	result := transport.ApiResponse("Success get all order data", http.StatusOK, constant.StatusSuccess, arrOrders)

	if len(arrOrders) == 0 {
		result = transport.ApiResponse("No orders yet", http.StatusOK, constant.StatusSuccess, len(arrOrders))
	}

	return &result
}
