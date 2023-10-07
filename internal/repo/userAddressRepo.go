package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func CreateUserAddress(userAddr *models.UserAddress) error {
	err := database.DB.Create(userAddr).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating user address")
		return helper.ErrCannotCreateEntity(helper.EntityUserAddress, err)
	}
	return nil
}

func GetUserAddress(condition map[string]interface{}) (*models.UserAddress, error) {
	var userAddr models.UserAddress
	if err := database.DB.Where(condition).First(&userAddr).Error; err != nil {
		log.Error().Err(err).Msg("Error fetching user address")
		return nil, helper.ErrCannotGetEntity(helper.EntityUserAddress, err)
	}
	return &userAddr, nil
}

func UpdateUserAddress(userId int, userAddr *models.UserAddressUpdate) error {
	_, err := GetUserAddress(map[string]interface{}{"user_id": userId})
	if err != nil {
		log.Error().Err(err).Msg("Error get user address for update")
		return err
	}

	err = database.DB.Where("user_id = ?", userId).Updates(userAddr).Error
	if err != nil {
		log.Error().Err(err).Msg("Error updating user address")
		return helper.ErrCannotUpdateEntity(helper.EntityUserAddress, err)
	}

	return nil
}
