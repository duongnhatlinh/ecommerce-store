package models

import (
	"time"
)

type UserAddress struct {
	Id         int       `json:"-" gorm:"id"`
	UserId     int       `json:"-" gorm:"user_id"`
	Address    string    `json:"address" gorm:"address" binding:"required"`
	City       string    `json:"city" gorm:"city" binding:"required"`
	Country    string    `json:"country" gorm:"country" binding:"required"`
	PostalCode string    `json:"postal_code" gorm:"postal_code" binding:"required"`
	Phone      string    `json:"phone" gorm:"phone" binding:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}

func (UserAddress) TableName() string {
	return "user_addresses"
}

type UserAddressUpdate struct {
	Address    string `json:"address" gorm:"address" binding:"required"`
	City       string `json:"city" gorm:"city" binding:"required"`
	Country    string `json:"country" gorm:"country" binding:"required"`
	PostalCode string `json:"postal_code" gorm:"postal_code" binding:"required"`
	Phone      string `json:"phone" gorm:"phone" binding:"required"`
}

func (UserAddressUpdate) TableName() string {
	return UserAddress{}.TableName()
}
