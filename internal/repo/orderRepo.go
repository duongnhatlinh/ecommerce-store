package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func CreateOrder(order *models.Order) (int, error) {
	err := database.DB.Create(order).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating order")
		return 0, helper.ErrCannotCreateEntity(helper.EntityOrder, err)
	}

	return order.Id, nil
}

func GetOrder(orderId, userId int) (*models.Order, error) {
	var order models.Order

	err := database.DB.Table(models.Order{}.TableName()).Where("id = ? and user_id = ?", orderId, userId).First(&order).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching order")
		return nil, helper.ErrCannotGetEntity(helper.EntityOrder, err)
	}

	if order.Status == 0 {
		log.Info().Msg("Order has been deleted")
		return nil, helper.ErrEntityDeleted(helper.EntityOrder)
	}
	return &order, nil
}

func GetListOrdersByUser(userId int) ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Table(models.Order{}.TableName()).Where("user_id = ? and status = ?", userId, 1).Find(&orders).Error
	if err != nil {
		log.Error().Err(err).Msg("Error listing orders by user")
		return nil, helper.ErrCannotListEntity(helper.EntityOrder, err)
	}
	return orders, nil
}

func GetTotalOrders(
	paging *helper.Paging,
) ([]models.Order, error) {
	var orders []models.Order
	db := database.DB.Table(models.Order{}.TableName()).Where("status = ?", 1)

	if err := db.Count(&paging.Total).Error; err != nil {
		log.Error().Err(err).Msg("Error counting total orders")
		return nil, helper.ErrDb(err)
	}

	db = db.Offset(paging.Limit * (paging.Page - 1))

	err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&orders).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching total orders")
		return nil, models.ErrCannotGetTotalEntity(helper.EntityOrder, err)
	}

	return orders, nil
}

func UpdateOrder(orderId int, totalPrice float32) error {
	err := database.DB.Table(models.Order{}.TableName()).Where("id = ?", orderId).Updates(map[string]interface{}{"total_price": totalPrice}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error updating order")
		return helper.ErrCannotUpdateEntity(helper.EntityOrder, err)
	}
	return nil
}

func DeleteOrder(orderId, userId int) error {
	_, err := GetOrder(orderId, userId)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching order for update")
		return err
	}
	err = database.DB.Table(models.Order{}.TableName()).Where("id = ?", orderId).Updates(map[string]interface{}{"status": 0}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error deleting order")
		return helper.ErrCannotDeleteEntity(helper.EntityOrder, err)
	}
	return nil
}
