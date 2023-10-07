package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func CreateInventory(inventory *models.Inventory) error {
	err := database.DB.Create(inventory).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating inventory")
		return helper.ErrCannotCreateEntity(helper.EntityInventory, err)
	}
	return nil
}

func GetInventory(condition map[string]interface{}) (*models.Inventory, error) {
	var inventory models.Inventory
	err := database.DB.Where(condition).First(&inventory).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching inventory")
		return nil, helper.ErrCannotGetEntity(helper.EntityInventory, err)
	}
	if inventory.Status == 0 {
		log.Info().Msg("Inventory has been deleted")
		return nil, helper.ErrEntityDeleted(helper.EntityInventory)
	}
	return &inventory, nil
}

func GetListInventories() ([]models.Inventory, error) {
	var inventories []models.Inventory
	if err := database.DB.Where("status = ?", 1).Find(&inventories).Error; err != nil {
		log.Error().Err(err).Msg("Error listing inventories")
		return nil, helper.ErrCannotListEntity(helper.EntityInventory, err)
	}
	return inventories, nil
}

func UpdateInventory(id int, inventory *models.Inventory) error {
	err := database.DB.Where("id = ? and status = ?", id, 1).Updates(inventory).Error
	if err != nil {
		log.Error().Err(err).Msg("Error updating inventory")
		return helper.ErrCannotUpdateEntity(helper.EntityInventory, err)
	}

	return nil
}

func DeleteInventory(id int) error {
	err := database.DB.Table(models.Inventory{}.TableName()).Where("id = ? and status = ?", id, 1).Updates(map[string]interface{}{"status": 0}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error deleting inventory")
		return helper.ErrCannotDeleteEntity(helper.EntityInventory, err)
	}

	return nil
}
