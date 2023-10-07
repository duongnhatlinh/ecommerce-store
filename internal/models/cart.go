package models

import (
	"ecommercestore/internal/helper"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Cart struct {
	UserId    int       `json:"user_id" gorm:"user_id"`
	ProductId int       `json:"product_id" gorm:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" gorm:"quantity" binding:"required"`
	Product   *Product  `json:"product" gorm:"foreignKey:ProductId;preload:false"`
	Status    int       `json:"status" gorm:"status;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

type ParamCart struct {
	ProductId int `form:"product_id"`
	Quantity  int `form:"quantity"`
}

func (Cart) TableName() string {
	return "carts"
}

var ErrItemOfCartNotExist = helper.NewCustomError(
	errors.New("item of cart does not exist"),
	"item of cart does not exist",
	"ErrItemOfCartNotExist",
)

func ErrCannotListItemsFromCart(entity string, err error) *helper.AppError {
	return helper.NewCustomError(
		err,
		fmt.Sprintf("Cannot list items from %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotListItemsFrom%s", entity),
	)
}
