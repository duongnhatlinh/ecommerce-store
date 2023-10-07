package testRepo

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDiscount(t *testing.T) {
	testSetup()
	discount, err := createMockDiscount()
	assert.NoError(t, err)
	assert.Equal(t, 1, discount.Id)
	assert.NotEmpty(t, discount.Desc)
	assert.NotEmpty(t, discount.Name)
	assert.Equal(t, 1, discount.Status)
	assert.NotEmpty(t, discount.Percent)
	assert.NotNil(t, discount.Active)
	assert.NotEmpty(t, discount.CreatedAt)
	assert.NotEmpty(t, discount.UpdatedAt)
}

func TestGetDiscount(t *testing.T) {
	testSetup()
	discount, err := createMockDiscount()
	assert.NoError(t, err)
	assert.Equal(t, 1, discount.Id)

	condition := map[string]interface{}{"id": discount.Id}
	fetchDiscount, err := repo.GetDiscount(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchDiscount)
	assert.Equal(t, discount.Id, fetchDiscount.Id)
	assert.Equal(t, discount.Name, fetchDiscount.Name)
	assert.Equal(t, discount.Desc, fetchDiscount.Desc)
	assert.Equal(t, discount.Percent, fetchDiscount.Percent)
	assert.Equal(t, discount.Active, fetchDiscount.Active)
}

func TestGetDiscountDeleted(t *testing.T) {
	testSetup()
	discount, err := createMockDiscount()
	assert.NoError(t, err)
	assert.Equal(t, 1, discount.Id)

	err = repo.DeleteDiscount(discount.Id)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": discount.Id}
	fetchDiscount, err := repo.GetDiscount(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchDiscount)
	assert.Equal(t, "ErrDISCOUNTDeleted", err.(*helper.AppError).Key)
}

func TestFetchNotExistingDiscount(t *testing.T) {
	testSetup()

	condition := map[string]interface{}{"id": 1}
	fetchDiscount, err := repo.GetDiscount(condition)
	assert.Error(t, err)
	assert.Nil(t, fetchDiscount)
	assert.Equal(t, "ErrCannotGetDISCOUNT", err.(*helper.AppError).Key)
}

func TestGetListDiscounts(t *testing.T) {
	testSetup()
	discount, err := createMockDiscount()
	assert.NoError(t, err)
	assert.Equal(t, 1, discount.Id)

	discount2, err := createMockDiscount2()
	assert.NoError(t, err)
	assert.Equal(t, 2, discount2.Id)

	discounts, err := repo.GetListDiscounts()
	assert.NoError(t, err)
	assert.NotEmpty(t, discounts)
	assert.Equal(t, 2, len(discounts))
}

func TestUpdateDiscount(t *testing.T) {
	testSetup()
	discount, err := createMockDiscount()
	assert.NoError(t, err)
	assert.Equal(t, 1, discount.Id)

	active := new(bool)
	*active = true
	data := &models.DiscountUpdate{
		Name:    "Holiday Special",
		Desc:    "Limited-time offer: 10% off holiday gifts",
		Percent: 10.0,
		Active:  active,
	}

	err = repo.UpdateDiscount(discount.Id, data)
	assert.NoError(t, err)

	condition := map[string]interface{}{"id": discount.Id}
	fetchDiscount, err := repo.GetDiscount(condition)
	assert.NoError(t, err)
	assert.NotNil(t, fetchDiscount)
	assert.Equal(t, data.Name, fetchDiscount.Name)
	assert.Equal(t, data.Desc, fetchDiscount.Desc)
	assert.Equal(t, data.Percent, fetchDiscount.Percent)
	assert.Equal(t, data.Active, fetchDiscount.Active)
}

func TestDeleteDiscount(t *testing.T) {
	testSetup()
	discount, err := createMockDiscount()
	assert.NoError(t, err)
	assert.Equal(t, 1, discount.Id)

	err = repo.DeleteDiscount(discount.Id)
	assert.NoError(t, err)
}
