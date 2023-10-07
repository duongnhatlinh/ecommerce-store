package repo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
)

func AddItemToOrder(userId int, orderItem *models.OrderItem) (*models.Order, error) {
	product, err := GetProduct(map[string]interface{}{"id": orderItem.ProductId}, "Discount")
	if err != nil {
		log.Error().Err(err).Msg("Error fetching product for action order")
		return nil, err
	}

	// discount
	orderItem.Price = product.Price
	percent := product.Discount.Percent
	if percent != 0 {
		orderItem.Price = orderItem.Price * (1 - percent/100)
	}

	order, err := GetOrder(orderItem.OrderId, userId)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching order for action order")
		return nil, err
	}

	// add item
	err = database.DB.Create(orderItem).Error
	if err != nil {
		log.Error().Err(err).Msg("Error adding product to order")
		return nil, helper.ErrCannotAddItemToEntity(helper.EntityOrderItem, err)
	}
	return order, nil
}

func RemoveItemFromOrder(itemId, orderId, userId int) error {
	orderItem, err := GetOrderItem(itemId)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching order item for action order")
		return err
	}

	order, err := GetOrder(orderId, userId)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching order for action order")
		return err
	}

	// remove item
	err = database.DB.
		Table(models.OrderItem{}.TableName()).
		Where("id = ?", itemId).
		Updates(map[string]interface{}{"status": 0}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error remove product from order")
		return helper.ErrCannotRemoveItemFromEntity(helper.EntityOrderItem, err)
	}

	// logic update order
	totalPrice := order.TotalPrice - orderItem.Price*float32(orderItem.Quantity)
	err = UpdateOrder(order.Id, totalPrice)
	if err != nil {
		log.Error().Err(err).Msg("Error update total price for order")
		return err
	}

	return nil
}

func GetListOrderItems(orderId int) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := database.DB.Where("order_id = ? and status = ?", orderId, 1).Find(&orderItems).Error
	if err != nil {
		log.Error().Err(err).Msg("Error listing order item")
		return nil, helper.ErrCannotListEntity(helper.EntityOrderItem, err)
	}

	return orderItems, nil
}

func UpdateOrderItem(itemId int, data *models.OrderItemUpdate) error {
	err := database.DB.Where("id = ?", itemId).Updates(data).Error
	if err != nil {
		log.Error().Err(err).Msg("Error update order item")
		return helper.ErrCannotUpdateEntity(helper.EntityOrderItem, err)
	}
	return nil
}

func GetOrderItem(itemId int) (*models.OrderItem, error) {
	var item models.OrderItem
	err := database.DB.
		Table(models.OrderItem{}.TableName()).
		Where("id = ?", itemId).
		First(&item).Error
	if err != nil {
		log.Error().Err(err).Msg("Error fetching order item")
		return nil, helper.ErrCannotGetEntity(helper.EntityOrderItem, err)
	}
	if item.Status == 0 {
		log.Info().Msg("Order item has been deleted")
		return nil, helper.ErrEntityDeleted(helper.EntityOrderItem)
	}

	return &item, nil
}
