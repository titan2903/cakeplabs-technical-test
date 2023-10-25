package entity

import (
	"time"

	"gorm.io/gorm"
)

type Addition struct {
	gorm.Model
	ID        int            `json:"id"`
	Name      string         `json:"name" gorm:"column:name;not null"`
	Price     float64        `json:"price" gorm:"column:price;not null"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted" form:"deleted"`
}
