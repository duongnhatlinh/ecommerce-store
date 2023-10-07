package models

import (
	"time"
)

type Inventory struct {
	Id        int       `json:"-" gorm:"id"`
	Quantity  int       `json:"quantity" gorm:"quantity" binding:"required"`
	Status    int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Inventory) TableName() string {
	return "inventories"
}
