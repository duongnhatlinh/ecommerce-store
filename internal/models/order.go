package models

import (
	"ecommercestore/internal/helper"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Order struct {
	Id         int     `json:"-" gorm:"id"`
	UserId     int     `json:"user_id" gorm:"user_id"`
	TotalPrice float32 `json:"total_price" gorm:"total_price" binding:"required"`
	//PaymentId int `json:"payment_id" gorm:"payment_id"`
	Status    int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

var ErrOrderUpdateValueInvalid = helper.NewCustomError(
	errors.New("update value order invalid"),
	"update value order invalid",
	"ErrOrderUpdateValueInvalid",
)

func ErrCannotGetTotalEntity(entity string, err error) *helper.AppError {
	return helper.NewCustomError(
		err,
		fmt.Sprintf("Cannot get total %ss", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGetTotal%sS", entity),
	)
}
