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

func TestSignUp(t *testing.T) {
	router := testSetup()
	body := userJson(models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Phone:     "1234567890",
		Password:  "password",
	})

	rec := performRequest(router, "POST", "/api/signup", body)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}

func TestSignUpEmptyFistName(t *testing.T) {
	router := testSetup()
	body := userJson(models.User{
		FirstName: "",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Phone:     "1234567890",
		Password:  "password",
	})

	rec := performRequest(router, "POST", "/api/signup", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "FirstName is required.", jsonFieldError(jsonResponse(rec.Body), "FirstName"))
}

func TestSignUpInvalidEmail(t *testing.T) {
	router := testSetup()
	body := userJson(models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoeexample.com",
		Phone:     "1234567890",
		Password:  "password",
	})

	rec := performRequest(router, "POST", "/api/signup", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Email must be in correct format.", jsonFieldError(jsonResponse(rec.Body), "FirstName"))
}

func TestSignIn(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	body := userSignInJson(models.UserSignIn{
		Email:    user.Email,
		Password: "password",
	})
	rec := performRequest(router, "POST", "/api/signin", body)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, jsonResponse(rec.Body)["data"])
}

func TestSignInInvalidEmail(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	body := userSignInJson(models.UserSignIn{
		Email:    "invalid" + user.Email,
		Password: "password",
	})
	rec := performRequest(router, "POST", "/api/signin", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, nil, jsonResponse(rec.Body)["data"])
	assert.Equal(t, "ErrEmailOrPasswordInvalid", jsonResponse(rec.Body)["error_key"])
}

func TestSignInInvalidPassword(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	body := userSignInJson(models.UserSignIn{
		Email:    user.Email,
		Password: "password_invalid",
	})
	rec := performRequest(router, "POST", "/api/signin", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, nil, jsonResponse(rec.Body)["data"])
	assert.Equal(t, "ErrEmailOrPasswordInvalid", jsonResponse(rec.Body)["error_key"])
}

func TestGetProfileUser(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	rec := performAuthorizedRequest(router, "GET", "/api/user/profile", "", token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, jsonResponse(rec.Body)["data"])
}

func TestUpdateInfoUser(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	body := userUpdateJson(models.InfoUser{
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "alicejohnson@example.com",
		Phone:     "5555555555",
	})
	rec := performAuthorizedRequest(router, "PUT", "/api/user/info", body, token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}

func TestUpdatePasswordUser(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	body := userPasswordJson(models.UserPassword{
		Password:    "password",
		NewPassword: "new_password",
	})

	rec := performAuthorizedRequest(router, "PUT", "/api/user/password", body, token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, jsonResponse(rec.Body)["data"])
}

func TestUpdatePasswordUserInvalidPassword(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	body := userPasswordJson(models.UserPassword{
		Password:    "invalid",
		NewPassword: "new_password",
	})

	rec := performAuthorizedRequest(router, "PUT", "/api/user/password", body, token)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "ErrPasswordInvalid", jsonResponse(rec.Body)["error_key"])
}

func TestGetUserById(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	rec := performAuthorizedRequest(router, "GET", fmt.Sprintf("/api/user/%d", user.Id), "", token)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, jsonResponse(rec.Body)["data"])
}

func TestGetUserByIdInvalid(t *testing.T) {
	router := testSetup()
	user := createMockUser(1)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user.Id, helper.Expiry)

	rec := performAuthorizedRequest(router, "GET", fmt.Sprintf("/api/user/%d", 10), "", token)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "ErrCannotGetUSER", jsonResponse(rec.Body)["error_key"])
}

func TestGetListUsers(t *testing.T) {
	router := testSetup()
	user1 := createMockUser(1)
	_ = createMockUser(2)
	provider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := provider.GenerateToken(user1.Id, helper.Expiry)

	rec := performAuthorizedRequestWithQueryParams(router, "GET", "/api/user/", "", token, 2, 3)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, jsonResponse(rec.Body)["data"])
}
