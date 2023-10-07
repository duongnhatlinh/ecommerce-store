package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	testSetup()
	category, err := createMockCategory()
	assert.NoError(t, err)
	assert.Equal(t, 1, category.Id)
	assert.NotEmpty(t, category.Desc)
	assert.NotEmpty(t, category.Name)
	assert.Equal(t, 1, category.Status)
	assert.NotEmpty(t, category.CreatedAt)
	assert.NotEmpty(t, category.UpdatedAt)
}

func TestGetCategory(t *testing.T) {
	testSetup()
	category, err := createMockCategory()
	assert.NoError(t, err)
	assert.Equal(t, 1, category.Id)

	condition := map[string]interface{}{"id": category.Id}
	fetchCategory, err := repo.GetCategory(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchCategory)
	assert.Equal(t, category.Id, fetchCategory.Id)
	assert.Equal(t, category.Name, fetchCategory.Name)
	assert.Equal(t, category.Desc, fetchCategory.Desc)
}

func TestGetCategoryDeleted(t *testing.T) {
	testSetup()
	category, err := createMockCategory()
	assert.NoError(t, err)
	assert.Equal(t, 1, category.Id)

	err = repo.DeleteCategory(category.Id)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": category.Id}
	fetchCategory, err := repo.GetCategory(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchCategory)
	assert.Equal(t, "ErrCATEGORYDeleted", err.(*helper.AppError).Key)
}

func TestFetchNotExistingCategory(t *testing.T) {
	testSetup()

	condition := map[string]interface{}{"id": 1}
	fetchCategory, err := repo.GetCategory(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchCategory)
	assert.Equal(t, "ErrCannotGetCATEGORY", err.(*helper.AppError).Key)
}

func TestGetListCategories(t *testing.T) {
	testSetup()
	category, err := createMockCategory()
	assert.NoError(t, err)
	assert.Equal(t, 1, category.Id)

	category2, err := createMockCategory2()
	assert.NoError(t, err)
	assert.Equal(t, 2, category2.Id)

	categories, err := repo.GetListCategories()
	assert.NoError(t, err)
	assert.NotEmpty(t, categories)
	assert.Equal(t, 2, len(categories))
}

func TestUpdateCategory(t *testing.T) {
	testSetup()
	category, err := createMockCategory()
	assert.NoError(t, err)
	assert.Equal(t, 1, category.Id)

	data := &models.CategoryUpdate{
		Name: "Men's Clothing3",
		Desc: "Category for men's clothing3.",
	}

	err = repo.UpdateCategory(category.Id, data)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": category.Id}
	fetchCategory, err := repo.GetCategory(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchCategory)
	assert.Equal(t, data.Desc, fetchCategory.Desc)
	assert.Equal(t, data.Name, fetchCategory.Name)
}

func TestDeleteCategory(t *testing.T) {
	testSetup()
	category, err := createMockCategory()
	assert.NoError(t, err)
	assert.Equal(t, 1, category.Id)

	err = repo.DeleteCategory(category.Id)
	assert.NoError(t, err)
}
