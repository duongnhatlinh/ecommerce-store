package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Salt)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}

func TestCreateUserWithExistingEmail(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	user, err = createMockUser()
	assert.Error(t, err)
	assert.Equal(t, "ErrEmailExisted", err.(*helper.AppError).Key)
}

func TestAuthenticate(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	data := &models.UserSignIn{
		Email:    user.Email,
		Password: "password",
	}

	token, err := repo.AuthenticateUser(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthenticateInvalidEmail(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	data := &models.UserSignIn{
		Email:    "invalid",
		Password: "password",
	}

	token, err := repo.AuthenticateUser(data)
	assert.Error(t, err)
	assert.Equal(t, "ErrEmailOrPasswordInvalid", err.(*helper.AppError).Key)
	assert.Empty(t, token)
}

func TestAuthenticateInvalidPassword(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	data := &models.UserSignIn{
		Email:    user.Email,
		Password: "invalid",
	}

	token, err := repo.AuthenticateUser(data)
	assert.Error(t, err)
	assert.Equal(t, "ErrEmailOrPasswordInvalid", err.(*helper.AppError).Key)
	assert.Empty(t, token)
}
func TestGetUserById(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	condition := map[string]interface{}{"id": user.Id}

	fetchedUser, err := repo.GetUser(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, user.Id, fetchedUser.Id)
	assert.Equal(t, user.FirstName, fetchedUser.FirstName)
	assert.Equal(t, user.LastName, fetchedUser.LastName)
	assert.Equal(t, user.Email, fetchedUser.Email)
	assert.Equal(t, user.Phone, fetchedUser.Phone)
	assert.Equal(t, user.Salt, fetchedUser.Salt)
	assert.Equal(t, user.Password, fetchedUser.Password)

}

func TestGetUserByEmail(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	condition := map[string]interface{}{"email": user.Email}

	fetchedUser, err := repo.GetUser(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, user.Id, fetchedUser.Id)
	assert.Equal(t, user.FirstName, fetchedUser.FirstName)
	assert.Equal(t, user.LastName, fetchedUser.LastName)
	assert.Equal(t, user.Email, fetchedUser.Email)
	assert.Equal(t, user.Phone, fetchedUser.Phone)
	assert.Equal(t, user.Salt, fetchedUser.Salt)
	assert.Equal(t, user.Password, fetchedUser.Password)

}

func TestUpdateInfoUser(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	data := &models.InfoUser{
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "janesmith@example.com",
		Phone:     "9876543210",
	}
	err = repo.UpdateInfoUser(user.Id, data)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": user.Id}
	fetchUser, err := repo.GetUser(condition)
	assert.NoError(t, err)
	assert.NotEmpty(t, fetchUser)
	assert.Equal(t, data.FirstName, fetchUser.FirstName)
	assert.Equal(t, data.LastName, fetchUser.LastName)
	assert.Equal(t, data.Email, fetchUser.Email)
	assert.Equal(t, data.Phone, fetchUser.Phone)
}

func TestUpdatePasswordUser(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	data := &models.UserPassword{
		Password:    "password",
		NewPassword: "newpassword",
	}

	err = repo.UpdatePasswordUser(user, data)
	assert.NoError(t, err)

	sigIn := &models.UserSignIn{
		Email:    user.Email,
		Password: "newpassword",
	}

	fetchUser, err := repo.AuthenticateUser(sigIn)
	assert.NoError(t, err)
	assert.NotEmpty(t, fetchUser)
}

func TestUpdatePasswordUserInvalidPassword(t *testing.T) {
	testSetup()
	user, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)

	data := &models.UserPassword{
		Password:    "invalid",
		NewPassword: "newpassword",
	}

	err = repo.UpdatePasswordUser(user, data)
	assert.Error(t, err)
	assert.Equal(t, "ErrPasswordInvalid", err.(*helper.AppError).Key)
}

func TestGetListUsers(t *testing.T) {
	testSetup()
	user1, err := createMockUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user1.Id)

	user2, err := createMockUser2()
	assert.NoError(t, err)
	assert.Equal(t, 2, user2.Id)

	paging := &helper.Paging{
		Page:  1,
		Limit: 10,
	}
	users, err := repo.GetListUsers(paging)
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Equal(t, int64(2), paging.Total)
}
