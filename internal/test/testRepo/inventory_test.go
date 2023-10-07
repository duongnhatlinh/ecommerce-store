package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateInventory(t *testing.T) {
	testSetup()
	inventory, err := createMockInventory()
	assert.NoError(t, err)
	assert.Equal(t, 1, inventory.Id)
	assert.NotEmpty(t, inventory.Quantity)
	assert.Equal(t, 1, inventory.Status)
	assert.NotEmpty(t, inventory.CreatedAt)
	assert.NotEmpty(t, inventory.UpdatedAt)
}

func TestGetInventory(t *testing.T) {
	testSetup()
	inventory, err := createMockInventory()
	assert.NoError(t, err)
	assert.Equal(t, 1, inventory.Id)

	condition := map[string]interface{}{"id": inventory.Id}
	fetchInventory, err := repo.GetInventory(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchInventory)
	assert.Equal(t, inventory.Id, fetchInventory.Id)
	assert.Equal(t, inventory.Quantity, fetchInventory.Quantity)
}

func TestGetInventoryDeleted(t *testing.T) {
	testSetup()
	inventory, err := createMockInventory()
	assert.NoError(t, err)
	assert.Equal(t, 1, inventory.Id)

	err = repo.DeleteInventory(inventory.Id)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": inventory.Id}
	fetchInventory, err := repo.GetInventory(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchInventory)
	assert.Equal(t, "ErrINVENTORYDeleted", err.(*helper.AppError).Key)
}

func TestFetchNotExistingInventory(t *testing.T) {
	testSetup()

	condition := map[string]interface{}{"id": 1}
	fetchInventory, err := repo.GetInventory(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchInventory)
	assert.Equal(t, "ErrCannotGetINVENTORY", err.(*helper.AppError).Key)
}

func TestGetListInventories(t *testing.T) {
	testSetup()
	inventory, err := createMockInventory()
	assert.NoError(t, err)
	assert.Equal(t, 1, inventory.Id)

	inventory2, err := createMockInventory2()
	assert.NoError(t, err)
	assert.Equal(t, 2, inventory2.Id)

	inventories, err := repo.GetListInventories()
	assert.NoError(t, err)
	assert.NotEmpty(t, inventories)
	assert.Equal(t, 2, len(inventories))
}

func TestUpdateInventory(t *testing.T) {
	testSetup()
	inventory, err := createMockInventory()
	assert.NoError(t, err)
	assert.Equal(t, 1, inventory.Id)

	data := &models.Inventory{
		Quantity: 50,
	}

	err = repo.UpdateInventory(inventory.Id, data)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": inventory.Id}
	fetchInventory, err := repo.GetInventory(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchInventory)
	assert.Equal(t, data.Quantity, fetchInventory.Quantity)
}

func TestDeleteInventory(t *testing.T) {
	testSetup()
	inventory, err := createMockInventory()
	assert.NoError(t, err)
	assert.Equal(t, 1, inventory.Id)

	err = repo.DeleteInventory(inventory.Id)
	assert.NoError(t, err)
}
