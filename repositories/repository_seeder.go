package repositories

import (
	"cakeplabs-technical-test/entity"
)

func (r *repositories) MenusSeeder() error {
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

	if err := r.db.Create(&menus).Error; err != nil {
		return err
	}

	return nil
}

func (r *repositories) AdditionsSeeder() error {
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

	if err := r.db.Create(&additions).Error; err != nil {
		return err
	}

	return nil
}
