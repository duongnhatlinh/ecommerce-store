package models

import (
	"time"
)

type OrderItem struct {
	Id        int       `json:"-" gorm:"id"`
	OrderId   int       `json:"order_id" gorm:"order_id"`
	ProductId int       `json:"product_id" gorm:"product_id" binding:"required"`
	Price     float32   `json:"price" gorm:"price"`
	Quantity  int       `json:"quantity" gorm:"quantity" binding:"required"`
	Status    int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

type OrderItemUpdate struct {
	Quantity int `json:"quantity" gorm:"quantity" binding:"required"`
}

func (OrderItemUpdate) TableName() string {
	return OrderItem{}.TableName()
}
