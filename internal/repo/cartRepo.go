package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func ItemOfCart(userId, productId int) (*models.Cart, error) {
	var item models.Cart
	db := database.DB.Table(models.Cart{}.TableName())
	err := db.
		Where("user_id = ? and product_id = ? and status = ?", userId, productId, 1).First(&item).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching item into cart")
		return nil, models.ErrItemOfCartNotExist
	}
	return &item, nil
}

func AddItemToCart(cart *models.Cart) error {
	db := database.DB.Table(models.Cart{}.TableName())

	item, err := ItemOfCart(cart.UserId, cart.ProductId)

	if err == nil {
		cart.Quantity = cart.Quantity + item.Quantity
		if _err := db.Where("user_id = ? and product_id = ?", cart.UserId, cart.ProductId).
			Updates(&cart).Error; _err != nil {
			log.Error().Err(err).Msg("Error adding product to cart")
			return helper.ErrCannotAddItemToEntity(helper.EntityCart, err)
		}
	} else {
		err = db.Create(cart).Error
		if err != nil {
			log.Error().Err(err).Msg("Error adding first product to cart")
			return helper.ErrCannotAddItemToEntity(helper.EntityCart, err)
		}
	}
	return nil
}

func RemoveItemFromCart(productId, userId int) error {
	db := database.DB.Table(models.Cart{}.TableName())

	_, err := ItemOfCart(userId, productId)
	if err != nil {
		log.Error().Err(err).Msg("Error verify product into cart")
		return err
	}
	err = db.Where("user_id = ? and product_id = ?", userId, productId).
		Updates(map[string]interface{}{"status": 0}).Error

	if err != nil {
		log.Error().Err(err).Msg("Error removing product from cart")
		return helper.ErrCannotRemoveItemFromEntity(helper.EntityCart, err)
	}
	return nil
}

func GetListItemsFromCart(userId int) (*[]models.Cart, error) {
	var cart []models.Cart
	db := database.DB.Table(models.Cart{}.TableName()).Where("status in (1)")

	db = db.Preload(helper.KeyProduct)

	err := db.Where("user_id", userId).Find(&cart).Error
	if err != nil {
		log.Error().Err(err).Msg("Error listing products to cart")
		return nil, models.ErrCannotListItemsFromCart(helper.EntityCart, err)
	}

	return &cart, nil
}
