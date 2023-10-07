package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	testSetup()
	order, err := createMockOrder()
	assert.NoError(t, err)
	assert.Equal(t, 1, order.Id)
	assert.Equal(t, 1, order.UserId)
	assert.Equal(t, 1, order.Status)
	assert.NotEmpty(t, order.CreatedAt)
	assert.NotEmpty(t, order.UpdatedAt)
}

func TestGetOrder(t *testing.T) {
	testSetup()
	order, err := createMockOrder()
	assert.NoError(t, err)
	assert.Equal(t, 1, order.Id)

	fetchOrder, err := repo.GetOrder(order.Id, order.UserId)
	assert.NoError(t, err)
	assert.NotNil(t, fetchOrder)
	assert.Equal(t, order.TotalPrice, fetchOrder.TotalPrice)
}

func TestFetchNotExistingOrder(t *testing.T) {
	testSetup()

	userId := 1
	orderId := 1
	fetchOrder, err := repo.GetOrder(orderId, userId)
	assert.Error(t, err)
	assert.Nil(t, fetchOrder)
	assert.Equal(t, "ErrCannotGetORDER", err.(*helper.AppError).Key)
}

func TestGetOrderDeleted(t *testing.T) {
	testSetup()
	order, err := createMockOrder()
	assert.NoError(t, err)
	assert.Equal(t, 1, order.Id)

	err = repo.DeleteOrder(order.Id, order.UserId)
	assert.NoError(t, err)

	fetchOrder, err := repo.GetOrder(order.Id, order.UserId)
	assert.Error(t, err)
	assert.Nil(t, fetchOrder)
	assert.Equal(t, "ErrORDERDeleted", err.(*helper.AppError).Key)
}

func TestGetListOrdersByUser(t *testing.T) {
	testSetup()
	userId, err := createMockListOrdersByUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, userId)

	orders, err := repo.GetListOrdersByUser(userId)
	assert.NoError(t, err)
	assert.NotNil(t, orders)
	assert.Equal(t, 2, len(orders))
}

func TestGetTotalOrders(t *testing.T) {
	testSetup()
	err := createMockListOrders()
	assert.NoError(t, err)

	paging := &helper.Paging{
		Page:  1,
		Limit: 10,
	}

	products, err := repo.GetTotalOrders(paging)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, int64(2), paging.Total)
}

func TestUpdateOrder(t *testing.T) {
	testSetup()
	order, err := createMockOrder()
	assert.NoError(t, err)
	assert.Equal(t, 1, order.Id)

	data := float32(30)

	err = repo.UpdateOrder(order.Id, data)
	assert.NoError(t, err)

	fetchOrder, err := repo.GetOrder(order.Id, order.UserId)
	assert.NoError(t, err)
	assert.NotNil(t, fetchOrder)
	assert.Equal(t, data, fetchOrder.TotalPrice)
}

func TestDeleteOrder(t *testing.T) {
	testSetup()
	order, err := createMockOrder()
	assert.NoError(t, err)
	assert.Equal(t, 1, order.Id)

	err = repo.DeleteOrder(order.Id, order.UserId)
	assert.NoError(t, err)
}
