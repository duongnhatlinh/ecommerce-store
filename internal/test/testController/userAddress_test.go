package testController

import (
	"ecommercestore/internal/conf"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/helper/jwt"
	"ecommercestore/internal/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateUserAddress(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)
	body := addressUserJson(models.UserAddress{
		Address:    "123 Main Street",
		City:       "New York",
		Country:    "USA",
		PostalCode: "10001",
		Phone:      "1234567890",
	})

	rec := performAuthorizedRequest(router, "POST", "/api/address/", body, token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}

func TestGetUserAddress(t *testing.T) {
	router := testSetup()
	addressUser := createMockUserAddress()
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(addressUser.UserId, helper.Expiry)

	rec := performAuthorizedRequest(router, "GET", "/api/address/", "", token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, jsonResponse(rec.Body)["data"])
}

func TestUpdateUserAddress(t *testing.T) {
	router := testSetup()
	addressUser := createMockUserAddress()
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(addressUser.UserId, helper.Expiry)
	body := userAddressUpdateJson(models.UserAddressUpdate{
		Address:    "234 Main Street",
		City:       "Eng",
		Country:    "England",
		PostalCode: "10341",
		Phone:      "1234567890",
	})

	rec := performAuthorizedRequest(router, "PUT", "/api/address/", body, token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}
