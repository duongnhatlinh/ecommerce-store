package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateActionLike(t *testing.T) {
	testSetup()
	proLike, err := createMockProductLike()
	assert.NoError(t, err)
	assert.NotNil(t, proLike)
	assert.NotEmpty(t, proLike.ProductId)
	assert.NotEmpty(t, proLike.UserId)
	assert.NotEmpty(t, proLike.CreatedAt)
	assert.NotEmpty(t, proLike.UpdatedAt)
}

func TestCreateActionDislike(t *testing.T) {
	testSetup()
	proLike, err := createMockProductLike()
	assert.NoError(t, err)
	assert.NotNil(t, proLike)

	err = repo.CreateActionDislike(proLike.ProductId, proLike.UserId)
	assert.NoError(t, err)
}

func TestCreateActionDislikeWhenHaveNotLike(t *testing.T) {
	testSetup()
	proLike, err := createMockProductLike()
	assert.NoError(t, err)
	assert.NotNil(t, proLike)

	err = repo.CreateActionDislike(proLike.ProductId, proLike.UserId)
	assert.NoError(t, err)

	err = repo.CreateActionDislike(proLike.ProductId, proLike.UserId)
	assert.Error(t, err)
	assert.Equal(t, "ErrCannotGetPRODUCT-LIKE", err.(*helper.AppError).Key)
}

func TestGetListUsersLike(t *testing.T) {
	testSetup()
	productId, err := createMockListUserLikeProduct()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, productId)

	paging := &helper.Paging{
		Page:  1,
		Limit: 10,
	}

	users, err := repo.GetListUsersLike(paging, productId)
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Equal(t, int64(2), paging.Total)
}
