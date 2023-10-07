package testRepo

import (
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserAddress(t *testing.T) {
	testSetup()
	userAddress, err := createMockUserAddress()
	assert.NoError(t, err)
	assert.Equal(t, 1, userAddress.Id)
	assert.NotEmpty(t, userAddress.CreatedAt)
	assert.NotEmpty(t, userAddress.UpdatedAt)
}

func TestGetUserAddress(t *testing.T) {
	testSetup()
	userAddress, err := createMockUserAddress()
	assert.NoError(t, err)
	assert.Equal(t, 1, userAddress.Id)

	condition := map[string]interface{}{"user_id": userAddress.UserId}
	fetchUserAddress, err := repo.GetUserAddress(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchUserAddress)
	assert.Equal(t, userAddress.Address, fetchUserAddress.Address)
	assert.Equal(t, userAddress.City, fetchUserAddress.City)
	assert.Equal(t, userAddress.Country, fetchUserAddress.Country)
	assert.Equal(t, userAddress.PostalCode, fetchUserAddress.PostalCode)
	assert.Equal(t, userAddress.Phone, fetchUserAddress.Phone)
}

func TestUpdateUserAddress(t *testing.T) {
	testSetup()
	userAddress, err := createMockUserAddress()
	assert.NoError(t, err)
	assert.Equal(t, 1, userAddress.Id)

	data := &models.UserAddressUpdate{
		Address:    "456 Elm Avenue",
		City:       "Los Angeles",
		Country:    "USA",
		PostalCode: "90001",
		Phone:      "9876543210",
	}

	err = repo.UpdateUserAddress(userAddress.UserId, data)
	assert.NoError(t, err)

	condition := map[string]interface{}{"user_id": userAddress.UserId}
	fetchUserAddress, err := repo.GetUserAddress(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchUserAddress)
	assert.Equal(t, data.Address, fetchUserAddress.Address)
	assert.Equal(t, data.City, fetchUserAddress.City)
	assert.Equal(t, data.Country, fetchUserAddress.Country)
	assert.Equal(t, data.PostalCode, fetchUserAddress.PostalCode)
	assert.Equal(t, data.Phone, fetchUserAddress.Phone)
}
