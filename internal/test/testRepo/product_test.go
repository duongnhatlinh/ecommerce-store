package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	testSetup()
	product, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)
	assert.Equal(t, 0, product.LikeCount)
	assert.Equal(t, 1, product.Status)
	assert.NotEmpty(t, product.CreatedAt)
	assert.NotEmpty(t, product.UpdatedAt)
}

func TestGetProduct(t *testing.T) {
	testSetup()
	product, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)

	condition := map[string]interface{}{"id": product.Id}
	fetchProduct, err := repo.GetProduct(condition, helper.KeyCategory, helper.KeyDiscount, helper.KeyInventory)
	assert.NoError(t, err)
	assert.NotNil(t, fetchProduct)
	assert.Equal(t, product.Name, fetchProduct.Name)
	assert.Equal(t, product.Code, fetchProduct.Code)
	assert.Equal(t, product.Color, fetchProduct.Color)
	assert.Equal(t, product.Size, fetchProduct.Size)
	assert.Equal(t, product.Desc, fetchProduct.Desc)
	assert.Equal(t, product.Price, fetchProduct.Price)
	assert.NotNil(t, fetchProduct.Category)
	assert.NotNil(t, fetchProduct.Discount)
	assert.NotNil(t, fetchProduct.Inventory)
}

func TestGetProductDeleted(t *testing.T) {
	testSetup()
	product, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)

	err = repo.DeleteProduct(product.Id)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": product.Id}
	fetchProduct, err := repo.GetProduct(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchProduct)
	assert.Equal(t, "ErrPRODUCTDeleted", err.(*helper.AppError).Key)
}

func TestFetchNotExistingProduct(t *testing.T) {
	testSetup()

	condition := map[string]interface{}{"id": 1}
	fetchProduct, err := repo.GetProduct(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchProduct)
	assert.Equal(t, "ErrCannotGetPRODUCT", err.(*helper.AppError).Key)
}

func TestGetListProductsWithFilter(t *testing.T) {
	testSetup()
	product1, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product1.Id)

	product2, err := createMockProduct2()
	assert.NoError(t, err)
	assert.Equal(t, 2, product2.Id)

	paging := &helper.Paging{
		Page:  1,
		Limit: 10,
	}

	filter := &models.FilterProduct{
		CategoryId: product1.CategoryId,
	}
	products, err := repo.GetListProducts(paging, filter, helper.KeyCategory, helper.KeyDiscount, helper.KeyInventory)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, int64(1), paging.Total)
}

func TestGetListProducts(t *testing.T) {
	testSetup()
	product1, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product1.Id)

	product2, err := createMockProduct2()
	assert.NoError(t, err)
	assert.Equal(t, 2, product2.Id)

	paging := &helper.Paging{
		Page:  1,
		Limit: 10,
	}

	products, err := repo.GetListProducts(paging, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, int64(2), paging.Total)
}

func TestUpdateProduct(t *testing.T) {
	testSetup()
	product, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)

	data := &models.ProductUpdate{
		Name:  "Product 3",
		Code:  "P0003",
		Color: "Green",
		Size:  12,
		Desc:  "Product 3 description.",
		Price: 23.99,
	}

	err = repo.UpdateProduct(product.Id, data)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": product.Id}
	fetchProduct, err := repo.GetProduct(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchProduct)
	assert.Equal(t, data.Name, fetchProduct.Name)
	assert.Equal(t, data.Code, fetchProduct.Code)
	assert.Equal(t, data.Color, fetchProduct.Color)
	assert.Equal(t, data.Size, fetchProduct.Size)
	assert.Equal(t, data.Desc, fetchProduct.Desc)
	assert.Equal(t, data.Price, fetchProduct.Price)
}

func TestDeleteProduct(t *testing.T) {
	testSetup()
	product, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)

	err = repo.DeleteProduct(product.Id)
	assert.NoError(t, err)
}

func TestIncreaseLikeCountProduct(t *testing.T) {
	testSetup()
	product, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)

	err = repo.IncreaseLikeCountProduct(product.Id)
	assert.NoError(t, err)
}

func TestDescendLikeCountProduct(t *testing.T) {
	testSetup()
	product, err := createMockProduct()
	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)

	err = repo.IncreaseLikeCountProduct(product.Id)
	assert.NoError(t, err)

	err = repo.DescendLikeCountProduct(product.Id)
	assert.NoError(t, err)
}
