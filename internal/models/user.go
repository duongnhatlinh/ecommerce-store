package models

import (
	"ecommercestore/internal/helper"
	"errors"
	"time"
)

type User struct {
	Id        int    `json:"-" gorm:"id"`
	FirstName string `json:"first_name" gorm:"first_name" binding:"required,min=1,max=50"`
	LastName  string `json:"last_name" gorm:"last_name" binding:"required,min=1,max=50"`
	Email     string `json:"email" gorm:"email" binding:"required,email"`
	Phone     string `json:"phone" gorm:"phone" binding:"required,min=10"`
	Password  string `json:"password" gorm:"password" binding:"required,min=5,max=50"`
	Salt      string `json:"-" gorm:"salt"`
	//Avatar
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type UserSignIn struct {
	Email    string `json:"email" gorm:"email" binding:"required,email"`
	Password string `json:"password" gorm:"password" binding:"required,min=5,max=50"`
}

type InfoUser struct {
	FirstName string `json:"first_name" gorm:"first_name" binding:"min=1,max=50"`
	LastName  string `json:"last_name" gorm:"last_name" binding:"min=1,max=50"`
	Email     string `json:"email" gorm:"email" binding:"email"`
	Phone     string `json:"phone" gorm:"phone" binding:"min=10"`
}

func (InfoUser) TableName() string {
	return User{}.TableName()
}

type UserPassword struct {
	Password    string `json:"password" gorm:"password" binding:"required,min=5,max=50"`
	NewPassword string `json:"new_password" gorm:"new_password" binding:"required,min=5,max=50"`
}

func (UserPassword) TableName() string {
	return User{}.TableName()
}

var (
	ErrEmailExisted = helper.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
	ErrEmailOrPasswordInvalid = helper.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)
	ErrPasswordInvalid = helper.NewCustomError(
		errors.New("password invalid"),
		"password invalid",
		"ErrPasswordInvalid",
	)
)
