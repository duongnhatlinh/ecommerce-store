package models

import "time"

type ProductLike struct {
	ProductId int       `json:"product_id" gorm:"product_id" binding:"required"`
	UserId    int       `json:"user_id" gorm:"user_id" binding:"required"`
	User      *User     `json:"user" gorm:"foreignKey:UserId;preload:false"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (ProductLike) TableName() string {
	return "product_likes"
}
