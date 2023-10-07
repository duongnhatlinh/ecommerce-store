package testController

import (
	"ecommercestore/internal/conf"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/helper/jwt"
	"ecommercestore/internal/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateDiscount(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	active := new(bool)
	*active = true
	body := discountJson(models.Discount{
		Name:    "Summer Sale",
		Desc:    "Get 20% off on all summer products",
		Percent: 20.0,
		Active:  active,
	})

	rec := performAuthorizedRequest(router, "POST", "/api/discount/", body, token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}

func TestGetDiscount(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	discount := createMockDiscount()
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	rec := performAuthorizedRequest(router, "GET", fmt.Sprintf("/api/discount/%d", discount.Id), "", token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, jsonResponse(rec.Body)["data"])
}

func TestGetListDiscounts(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	_ = createMockDiscount()
	_ = createMockDiscount()
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	rec := performAuthorizedRequest(router, "GET", "/api/discount/", "", token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, jsonResponse(rec.Body)["data"])
}

func TestUpdateDiscount(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	discount := createMockDiscount()
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	active := new(bool)
	*active = false
	body := discountUpdateJson(models.DiscountUpdate{
		Name:    "Back-to-School Sale",
		Desc:    "Save 15% on back-to-school essentials",
		Percent: 15.0,
		Active:  active,
	})

	rec := performAuthorizedRequest(router, "PUT", fmt.Sprintf("/api/discount/%d", discount.Id), body, token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}

func TestDeleteDiscount(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	discount := createMockDiscount()
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	rec := performAuthorizedRequest(router, "DELETE", fmt.Sprintf("/api/discount/%d", discount.Id), "", token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}
