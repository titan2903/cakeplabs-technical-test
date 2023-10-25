package repositories

import (
	"cakeplabs-technical-test/entity"

	"gorm.io/gorm"
)

type repositories struct {
	db *gorm.DB
}

type Repositories interface {
	GetMenuById(id int) (entity.Menu, error)
	GetAllMenusByKeyword(keywords string) []entity.Menu
	GetAllMenus() []entity.Menu

	GetAllOrdersByKeyword(keywords string) []entity.Order
	GetOrderById(id int) (entity.Order, error)
	GetAllOrders() []entity.Order
	CreateOrder(order entity.Order) error

	GetAllAdditionsByKeyword(keywords string) []entity.Addition
	GetAllAdditions() []entity.Addition
	GetAdditionById(id int) (entity.Addition, error)

	MenusSeeder() error
	AdditionsSeeder() error
}

func NewRepositories(
	q *gorm.DB,
	// mg *mongo.Client,
) Repositories {
	return &repositories{
		db: q,
		// mg: mg,
	}
}
