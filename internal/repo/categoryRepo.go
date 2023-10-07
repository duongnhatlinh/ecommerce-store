package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func CreateCategory(category *models.Category) error {
	err := database.DB.Create(category).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating category")
		return helper.ErrCannotCreateEntity(helper.EntityCategory, err)
	}
	return nil
}

func GetCategory(condition map[string]interface{}) (*models.Category, error) {
	var category models.Category
	err := database.DB.Where(condition).First(&category).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching category")
		return nil, helper.ErrCannotGetEntity(helper.EntityCategory, err)
	}

	if category.Status == 0 {
		log.Info().Msg("Category has been deleted")
		return nil, helper.ErrEntityDeleted(helper.EntityCategory)
	}
	return &category, nil
}

func GetListCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Where("status= ?", 1).Find(&categories).Error; err != nil {
		log.Error().Err(err).Msg("Error listing categories")
		return nil, helper.ErrCannotListEntity(helper.EntityCategory, err)
	}
	return categories, nil
}

func UpdateCategory(id int, category *models.CategoryUpdate) error {
	err := database.DB.Where("id =? and status = ?", id, 1).Updates(category).Error
	if err != nil {
		log.Error().Err(err).Msg("Error updating category")
		return helper.ErrCannotUpdateEntity(helper.EntityCategory, err)
	}

	return nil
}

func DeleteCategory(id int) error {
	err := database.DB.Table(models.Category{}.TableName()).Where("id= ? and status = ?", id, 1).Updates(map[string]interface{}{"status": 0}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error deleting category")
		return helper.ErrCannotDeleteEntity(helper.EntityCategory, err)
	}

	return nil
}
