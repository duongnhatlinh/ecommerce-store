package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddItemToOrder(t *testing.T) {
	testSetup()
	orderItem, order, err := createMockOrderItem()
	assert.NoError(t, err)
	assert.NotNil(t, orderItem)
	assert.NotNil(t, order)
	assert.NotEmpty(t, orderItem.OrderId)
	assert.NotEmpty(t, orderItem.ProductId)
	assert.NotEmpty(t, orderItem.Price)
	assert.NotEmpty(t, orderItem.Quantity)
	assert.Equal(t, 1, orderItem.Status)
	assert.NotEmpty(t, orderItem.CreatedAt)
	assert.NotEmpty(t, orderItem.UpdatedAt)
}

func TestRemoveItemFromOrder(t *testing.T) {
	testSetup()
	orderItem, order, err := createMockOrderItem()
	assert.NoError(t, err)
	assert.NotNil(t, orderItem)
	assert.NotNil(t, order)

	err = repo.RemoveItemFromOrder(orderItem.ProductId, orderItem.OrderId, order.UserId)
	assert.NoError(t, err)

	fetchOrderItem, err := repo.GetOrderItem(orderItem.Id)
	assert.Error(t, err)
	assert.Nil(t, fetchOrderItem)
	assert.Equal(t, "ErrORDER-ITEMDeleted", err.(*helper.AppError).Key)
}

func TestGetListOrderItems(t *testing.T) {
	testSetup()
	orderId, err := createMockListOrderItems()
	assert.NoError(t, err)
	assert.Equal(t, 1, orderId)

	orderItems, err := repo.GetListOrderItems(orderId)
	assert.NoError(t, err)
	assert.NotNil(t, orderItems)
	assert.Equal(t, 2, len(orderItems))
}

func TestGetOrderItem(t *testing.T) {
	testSetup()
	orderItem, order, err := createMockOrderItem()
	assert.NoError(t, err)
	assert.NotNil(t, orderItem)
	assert.NotNil(t, order)

	fetchOrderItem, err := repo.GetOrderItem(orderItem.Id)
	assert.NoError(t, err)
	assert.NotNil(t, fetchOrderItem)
	assert.Equal(t, orderItem.Price, fetchOrderItem.Price)
	assert.Equal(t, orderItem.Quantity, fetchOrderItem.Quantity)
}
