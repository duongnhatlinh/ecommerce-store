package models

import (
	"ecommercestore/internal/helper"
	"fmt"
	"time"
)

type Product struct {
	Id          int        `json:"-" gorm:"id"`
	Name        string     `json:"name" gorm:"name" binding:"required,min=1,max=200"`
	Code        string     `json:"code" gorm:"code" binding:"required,len=5"`
	Color       string     `json:"color" gorm:"color" binding:"required"`
	Size        int        `json:"size" gorm:"size" binding:"required"`
	Desc        string     `json:"desc" gorm:"desc" binding:"required,min=5,max=1000"`
	Price       float32    `json:"price" gorm:"price" binding:"required"`
	CategoryId  int        `json:"-" gorm:"category_id"`
	Category    *Category  `json:"category,omitempty" gorm:"foreignKey:CategoryId;preload:false"`
	InventoryId int        `json:"-" gorm:"inventory_id"`
	Inventory   *Inventory `json:"inventory,omitempty" gorm:"foreignKey:InventoryId;preload:false"`
	DiscountId  int        `json:"-" gorm:"discount_id"`
	Discount    *Discount  `json:"discount,omitempty" gorm:"foreignKey:DiscountId;preload:false"`
	//Images
	LikeCount int       `json:"like_count"gorm:"column:liked_count"` // computed field
	Status    int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}

type ProductUpdate struct {
	Name  string  `json:"name" gorm:"name" binding:"min=1,max=200"`
	Code  string  `json:"code" gorm:"code" binding:"min=5,max=20"`
	Color string  `json:"color" gorm:"color"`
	Size  int     `json:"size" gorm:"size" `
	Desc  string  `json:"desc" gorm:"desc" binding:"min=5,max=1000"`
	Price float32 `json:"price" gorm:"price"`
	//Images
	Status int `json:"status" gorm:"column:status;default:1"`
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}

type ProductForeignKey struct {
	CategoryId  int `json:"category_id" form:"category_id"`
	InventoryId int `json:"inventory_id" form:"inventory_id"`
	DiscountId  int `json:"discount_id" form:"discount_id"`
}

func ErrCannotIncreaseLikeCountProduct(err error) *helper.AppError {
	return helper.NewCustomError(
		err,
		fmt.Sprintf("Cannot increase likecount product"),
		fmt.Sprintf("ErrCannotIncreaseLikeCountProduct"),
	)
}

func ErrCannotDescendLikeCountProduct(err error) *helper.AppError {
	return helper.NewCustomError(
		err,
		fmt.Sprintf("Cannot descennd likecount product"),
		fmt.Sprintf("ErrCannotDescendLikeCountProduct"),
	)
}

type FilterProduct struct {
	CategoryId int `json:"category_id,omitempty" form:"category_id"`
}
