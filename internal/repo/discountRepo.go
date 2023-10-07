package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func CreateDiscount(discount *models.Discount) error {
	err := database.DB.Create(discount).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating discount")
		return helper.ErrCannotCreateEntity(helper.EntityDiscount, err)
	}
	return nil
}

func GetDiscount(condition map[string]interface{}) (*models.Discount, error) {
	var discount models.Discount
	err := database.DB.Table(models.Discount{}.TableName()).Where(condition).First(&discount).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching discount")
		return nil, helper.ErrCannotGetEntity(helper.EntityDiscount, err)
	}

	if discount.Status == 0 {
		log.Info().Msg("Discount has been deleted")
		return nil, helper.ErrEntityDeleted(helper.EntityDiscount)
	}
	return &discount, nil
}

func GetListDiscounts() ([]models.Discount, error) {
	var discounts []models.Discount
	if err := database.DB.Where("status in (1)").Find(&discounts).Error; err != nil {
		log.Error().Err(err).Msg("Error listing discounts")
		return nil, helper.ErrCannotListEntity(helper.EntityDiscount, err)
	}
	return discounts, nil
}

func UpdateDiscount(id int, discount *models.DiscountUpdate) error {
	err := database.DB.Where("id = ? and status = ?", id, 1).Updates(discount).Error
	if err != nil {
		log.Error().Err(err).Msg("Error updating discount")
		return helper.ErrCannotUpdateEntity(helper.EntityDiscount, err)
	}

	return nil
}

func DeleteDiscount(id int) error {
	err := database.DB.Table(models.Discount{}.TableName()).Where("id = ? and status = ?", id, 1).Updates(map[string]interface{}{"status": 0}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error deleting discount")
		return helper.ErrCannotDeleteEntity(helper.EntityDiscount, err)
	}

	return nil
}
