package controllers

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		product := ctx.MustGet(gin.BindKey).(*models.Product)

		var params models.ProductForeignKey
		if err := ctx.BindQuery(&params); err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		product.CategoryId = params.CategoryId
		product.InventoryId = params.InventoryId
		product.DiscountId = params.DiscountId

		err := repo.CreateProduct(product)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}
		product, err := repo.GetProduct(map[string]interface{}{"id": id}, helper.KeyCategory, helper.KeyDiscount, helper.KeyInventory)
		if err != nil {
			panic(err)

		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(product))
	}
}
func GetListProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging helper.Paging

		err := ctx.BindQuery(&paging)
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		var filter models.FilterProduct
		err = ctx.BindQuery(&filter)
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		products, err := repo.GetListProducts(&paging, &filter, helper.KeyCategory, helper.KeyDiscount, helper.KeyInventory)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.NewSuccessRespond(products, paging, filter))
	}
}
func UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		product := ctx.MustGet(gin.BindKey).(*models.ProductUpdate)

		err = repo.UpdateProduct(id, product)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
func DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		err = repo.DeleteProduct(id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
