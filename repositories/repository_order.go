package repositories

import (
	"cakeplabs-technical-test/entity"

	"github.com/labstack/gommon/log"
)

func (r *repositories) GetAllOrdersByKeyword(keywords string) []entity.Order {
	var orders []entity.Order
	result := r.db.Preload("Menu").Joins("JOIN menus ON orders.menu_id = menus.id").
		Where("orders.topping_name LIKE ? OR orders.filling_name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").
		Find(&orders)

	if result.Error != nil {
		log.Error("Error while fetching orders:", result.Error)
		return nil
	}

	return orders
}

func (r *repositories) GetOrderById(id int) (entity.Order, error) {
	var order entity.Order
	result := r.db.Preload("Menu").Where("id = ?", id).First(&order)
	return order, result.Error
}

func (r *repositories) GetAllOrders() []entity.Order {
	var orders []entity.Order
	_ = r.db.Find(&orders)

	return orders
}

func (r *repositories) CreateOrder(order entity.Order) error {
	if err := r.db.Create(&order).Error; err != nil {
		return err
	}

	return nil
}
