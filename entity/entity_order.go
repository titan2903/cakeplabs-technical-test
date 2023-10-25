package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID          int            `json:"id" gorm:"column:id;not null"`
	TotalPrice  float64        `json:"total_price" gorm:"column:total_price;not null" validate:"required"`
	MenuID      int            `json:"menu_id" gorm:"column:menu_id;not null" validate:"required"`
	ToppingName string         `json:"topping_name" gorm:"column:topping_name"`
	FillingName string         `json:"filling_name" gorm:"column:filling_name"`
	CreatedAt   time.Time      `json:"created_at" form:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time      `json:"updated_at" form:"updated_at" gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted" form:"deleted"`
	Menu        Menu           `gorm:"foreignKey:MenuID"`
}
