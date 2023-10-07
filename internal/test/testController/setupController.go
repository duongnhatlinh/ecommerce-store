package testController

import (
	"bytes"
	"ecommercestore/internal/database"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"ecommercestore/internal/routes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

func testSetup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.ResetTestDatabase()

	router := routes.SetRoute()
	return router
}

func jsonResponse(body *bytes.Buffer) map[string]interface{} {
	data := map[string]interface{}{}
	if err := json.Unmarshal(body.Bytes(), &data); err != nil {
		log.Panic().Err(err).Msg("Error unmarshalling JSON body.")
	}
	return data
}

func jsonFieldError(jsonRes map[string]interface{}, field string) interface{} {
	data, ok := jsonRes["error"].(map[string]interface{})
	if !ok {
		log.Panic().Interface("jsonRes", jsonRes).Msg("JSON response data is not a map.")
	}
	return data[field]
}

func NewRequest(method, path, body string) *http.Request {
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		log.Panic().Err(err).Msg("Error creating new request")
	}

	req.Header.Set("Content-Type", "application/json")
	return req
}

func performRequest(router *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
	req := NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func performAuthorizedRequest(router *gin.Engine, method, path, body, token string) *httptest.ResponseRecorder {
	req := NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+token)
	router.ServeHTTP(rec, req)
	return rec
}

func performAuthorizedRequestWithQueryParams(router *gin.Engine, method, path, body, token string, page, limit int) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		log.Panic().Err(err).Msg("Error creating new request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.URL.Query().Set("page", strconv.Itoa(page))
	req.URL.Query().Set("limit", strconv.Itoa(limit))
	rec := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+token)
	router.ServeHTTP(rec, req)
	return rec
}

// user

func createMockUser(varity int) *models.User {
	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     fmt.Sprintf("johndoe%d@example.com", varity),
		Phone:     "1234567890",
		Password:  "password",
	}
	err := repo.CreateUser(user)
	if err != nil {
		log.Panic().Err(err).Msg("Error creating test user.")
	}
	return user
}

func userJson(user models.User) string {
	body, err := json.Marshal(map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"phone":      user.Phone,
		"password":   user.Password,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

// userSignIn

func userSignInJson(user models.UserSignIn) string {
	body, err := json.Marshal(map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

// user update

func userUpdateJson(user models.InfoUser) string {
	body, err := json.Marshal(map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"phone":      user.Phone,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

// user update

func userPasswordJson(user models.UserPassword) string {
	body, err := json.Marshal(map[string]interface{}{
		"password":     user.Password,
		"new_password": user.NewPassword,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

// userAddress

func createMockUserAddress() *models.UserAddress {
	user := createMockUser(1)
	userId := user.Id
	userAddress := &models.UserAddress{
		UserId:     userId,
		Address:    "123 Main Street",
		City:       "New York",
		Country:    "USA",
		PostalCode: "10001",
		Phone:      "1234567890",
	}

	err := repo.CreateUserAddress(userAddress)
	if err != nil {
		log.Panic().Err(err).Msg("Error creating test user.")
	}
	return userAddress
}

func addressUserJson(userAddr models.UserAddress) string {
	body, err := json.Marshal(map[string]interface{}{
		"address":     userAddr.Address,
		"city":        userAddr.City,
		"country":     userAddr.Country,
		"postal_code": userAddr.PostalCode,
		"phone":       userAddr.Phone,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

func userAddressUpdateJson(userAddr models.UserAddressUpdate) string {
	body, err := json.Marshal(map[string]interface{}{
		"address":     userAddr.Address,
		"city":        userAddr.City,
		"country":     userAddr.Country,
		"postal_code": userAddr.PostalCode,
		"phone":       userAddr.Phone,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

// discount
func createMockDiscount() *models.Discount {
	active := new(bool)
	*active = true
	discount := &models.Discount{
		Name:    "Summer Sale",
		Desc:    "Get 20% off on all summer products",
		Percent: 20.0,
		Active:  active,
	}

	err := repo.CreateDiscount(discount)
	if err != nil {
		log.Panic().Err(err).Msg("Error creating test user.")
	}
	return discount
}

func discountJson(discount models.Discount) string {
	body, err := json.Marshal(map[string]interface{}{
		"name":    discount.Name,
		"desc":    discount.Desc,
		"percent": discount.Percent,
		"active":  discount.Active,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

func discountUpdateJson(discount models.DiscountUpdate) string {
	body, err := json.Marshal(map[string]interface{}{
		"name":    discount.Name,
		"desc":    discount.Desc,
		"percent": discount.Percent,
		"active":  discount.Active,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}