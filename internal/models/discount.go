package models

import (
	"time"
)

type Discount struct {
	Id        int       `json:"-" gorm:"id"`
	Name      string    `json:"name" gorm:"name" binding:"required,min=1,max=200"`
	Desc      string    `json:"desc" gorm:"desc" binding:"required,min=5,max=1000"`
	Percent   float32   `json:"percent" gorm:"percent" binding:"required,lte=100,gte=1"`
	Active    *bool     `json:"active" gorm:"active" binding:"required,boolean"`
	Status    int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Discount) TableName() string {
	return "discounts"
}

type DiscountUpdate struct {
	Name    string  `json:"name" gorm:"name" binding:"min=1,max=200"`
	Desc    string  `json:"desc" gorm:"desc" binding:"min=5,max=1000"`
	Percent float32 `json:"percent" gorm:"percent"  binding:"required"`
	Active  *bool   `json:"active" gorm:"active"  binding:"required"`
}

func (DiscountUpdate) TableName() string {
	return Discount{}.TableName()
}
