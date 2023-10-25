package seed

import (
	"cakeplabs-technical-test/database/config"
	"cakeplabs-technical-test/entity"
	"log"

	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeed() *seed {
	return &seed{config.GetQuery()}
}

func (s *seed) MenusSeeder() {
	var menus = []entity.Menu{
		{
			Name:  "Pizza",
			Price: 50000,
		},
		{
			Name:  "Doughnut",
			Price: 20000,
		},
		{
			Name:  "Pie",
			Price: 45000,
		},
	}

	if err := s.DB.Create(&menus).Error; err != nil {
		log.Printf("cannot seed data menus, with error %v\n", err)
	}
	log.Println("success seed data menus")
}

func (s *seed) AdditionsSeeder() {
	var additions = []entity.Addition{
		{
			Name:  "Cheese",
			Price: 12000,
		},
		{
			Name:  "Chicken",
			Price: 18000,
		},
		{
			Name:  "Pepper",
			Price: 8000,
		},
		{
			Name:  "Tomato",
			Price: 9000,
		},
		{
			Name:  "Tuna",
			Price: 20000,
		},
		{
			Name:  "Blueberry",
			Price: 12000,
		},
		{
			Name:  "Sugar Glaze",
			Price: 10000,
		},
		{
			Name:  "Apple Slices",
			Price: 14000,
		},
		{
			Name:  "Milk Cream",
			Price: 10000,
		},
	}

	if err := s.DB.Create(&additions).Error; err != nil {
		log.Printf("cannot seed data Addition, with error %v\n", err)
	}
	log.Println("success seed data addition")
}
