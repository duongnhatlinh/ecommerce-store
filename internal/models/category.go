package models

import (
	"time"
)

type Category struct {
	Id   int    `json:"-" gorm:"id"`
	Name string `json:"name" gorm:"name" binding:"required,min=1,max=100"`
	Desc string `json:"desc" gorm:"desc" binding:"required,min=5,max=1000"`
	//Icon
	Status    int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryUpdate struct {
	Name string `json:"name" gorm:"name" binding:"min=1,max=100"`
	Desc string `json:"desc" gorm:"desc" binding:"min=5,max=1000"`
	//Icon
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}
