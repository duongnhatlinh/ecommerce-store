package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddItemToCartEmpty(t *testing.T) {
	testSetup()
	cart, err := createMockCart()
	assert.NoError(t, err)
	assert.NotNil(t, cart)
	assert.NotEmpty(t, cart.UserId)
	assert.NotEmpty(t, cart.ProductId)
	assert.NotEmpty(t, cart.Quantity)
	assert.Equal(t, 1, cart.Status)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.UpdatedAt)
}

func TestAddItemToCartExistedThatItem(t *testing.T) {
	testSetup()
	cart1, err := createMockCart()
	assert.NoError(t, err)
	assert.NotNil(t, cart1)

	cart2, err := createMockCart()
	assert.NoError(t, err)
	assert.NotNil(t, cart2)

	assert.Equal(t, cart1.Quantity, cart2.Quantity)
}

func TestItemOfCart(t *testing.T) {
	testSetup()
	cart, err := createMockCart()
	assert.NoError(t, err)
	assert.NotNil(t, cart)

	fetchCart, err := repo.ItemOfCart(cart.UserId, cart.ProductId)
	assert.NoError(t, err)
	assert.NotNil(t, fetchCart)
}

func TestNotItemOfCart(t *testing.T) {
	testSetup()
	cart, err := createMockCart()
	assert.NoError(t, err)
	assert.NotNil(t, cart)

	fetchCart, err := repo.ItemOfCart(cart.UserId, 2)
	assert.Error(t, err)
	assert.Nil(t, fetchCart)
	assert.Equal(t, "ErrItemOfCartNotExist", err.(*helper.AppError).Key)
}

func TestRemoveItemFromCart(t *testing.T) {
	testSetup()
	cart, err := createMockCart()
	assert.NoError(t, err)
	assert.NotNil(t, cart)

	err = repo.RemoveItemFromCart(cart.ProductId, cart.UserId)
	assert.NoError(t, err)

	fetchCart, err := repo.ItemOfCart(cart.UserId, cart.ProductId)
	assert.Error(t, err)
	assert.Nil(t, fetchCart)
}

func TestGetListItemsFromCart(t *testing.T) {
	testSetup()
	userId, err := createMockListItemsOfCart()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)

	carts, err := repo.GetListItemsFromCart(userId)
	assert.NoError(t, err)
	assert.NotNil(t, carts)
	assert.Equal(t, 2, len(*carts))
}
