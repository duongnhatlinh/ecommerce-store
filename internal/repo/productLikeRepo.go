package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func CreateActionLike(proLike *models.ProductLike) error {
	err := database.DB.Create(proLike).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating action like for product")
		return helper.ErrCannotCreateEntity(helper.EntityProductLike, err)
	}

	err = IncreaseLikeCountProduct(proLike.ProductId)
	if err != nil {
		log.Error().Err(err).Msg("Error increasing like product")
		return models.ErrCannotIncreaseLikeCountProduct(err)
	}
	return nil
}

func CreateActionDislike(productId, userId int) error {
	var proLike models.ProductLike

	if err := database.DB.Table(models.ProductLike{}.TableName()).
		Where("user_id = ? and product_id = ?", userId, productId).
		First(&proLike).Error; err != nil {
		log.Error().Err(err).Msg("Error fetching product like for action dislike ")
		return helper.ErrCannotGetEntity(helper.EntityProductLike, err)
	}

	if err := database.DB.Table(models.ProductLike{}.TableName()).
		Where("user_id = ? and product_id = ?", userId, productId).
		Delete(nil).Error; err != nil {
		log.Error().Err(err).Msg("Error deleting product like for action dislike ")
		return helper.ErrCannotDeleteEntity(helper.EntityProductLike, err)
	}

	err := DescendLikeCountProduct(productId)
	if err != nil {
		log.Error().Err(err).Msg("Error descending like product")
		return models.ErrCannotDescendLikeCountProduct(err)
	}
	return nil
}

func GetListUsersLike(
	paging *helper.Paging,
	productId int,
) ([]models.ProductLike, error) {
	var usersLike []models.ProductLike

	db := database.DB
	db = db.Table(models.ProductLike{}.TableName()).Where("product_id = ?", productId)

	if err := db.Count(&paging.Total).Error; err != nil {
		log.Error().Err(err).Msg("Error counting user like product")
		return nil, helper.ErrDb(err)
	}

	db = db.Preload(helper.KeyUser)

	db = db.Offset(paging.Limit * (paging.Page - 1))

	err := db.
		Table(models.ProductLike{}.TableName()).
		Limit(paging.Limit).
		Order("product_id desc").
		Find(&usersLike).Error
	if err != nil {
		log.Error().Err(err).Msg("Error listing users like product")
		return nil, helper.ErrCannotListEntity(helper.EntityProductLike, err)
	}

	return usersLike, nil
}
