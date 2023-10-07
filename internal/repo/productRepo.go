package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateProduct(product *models.Product) error {
	err := database.DB.Create(product).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating product")
		return helper.ErrCannotCreateEntity(helper.EntityProduct, err)
	}
	return nil
}

func GetProduct(condition map[string]interface{}, moreKeys ...string) (*models.Product, error) {
	var product models.Product
	db := database.DB

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	err := db.Where(condition).First(&product).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching product")
		return nil, helper.ErrCannotGetEntity(helper.EntityProduct, err)
	}

	if product.Status == 0 {
		log.Info().Msg("Product has been deleted")
		return nil, helper.ErrEntityDeleted(helper.EntityProduct)
	}
	return &product, nil
}

func GetListProducts(paging *helper.Paging,
	filter *models.FilterProduct,
	moreKeys ...string,
) ([]models.Product, error) {
	var products []models.Product

	db := database.DB.Table(models.Product{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if filter.CategoryId > 0 {
			db = db.Where("category_id = ?", filter.CategoryId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		log.Error().Err(err).Msg("Error counting product")
		return nil, helper.ErrDb(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Offset(paging.Limit * (paging.Page - 1))

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&products).
		Error; err != nil {
		log.Error().Err(err).Msg("Error listing user")
		return nil, helper.ErrCannotListEntity(helper.EntityProduct, err)
	}

	return products, nil
}

func UpdateProduct(id int, product *models.ProductUpdate) error {
	err := database.DB.Where("id = ? and status = ?", id, 1).Updates(product).Error
	if err != nil {
		log.Error().Err(err).Msg("Error update product")
		return helper.ErrCannotUpdateEntity(helper.EntityProduct, err)
	}

	return nil
}

func DeleteProduct(id int) error {
	err := database.DB.Table(models.Product{}.TableName()).Where("id = ? and status = ?", id, 1).Updates(map[string]interface{}{"status": 0}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error deleting product")
		return helper.ErrCannotDeleteEntity(helper.EntityProduct, err)
	}

	return nil
}

func IncreaseLikeCountProduct(id int) error {
	db := database.DB

	if err := db.Table(models.Product{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		log.Error().Err(err).Msg("Error increasing like count for product ")
		return helper.ErrDb(err)
	}

	return nil
}

func DescendLikeCountProduct(id int) error {
	db := database.DB

	if err := db.Table(models.Product{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		log.Error().Err(err).Msg("Error descending like count for product ")
		return helper.ErrDb(err)
	}

	return nil
}
