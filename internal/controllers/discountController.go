package controllers

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateDiscount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		discount := ctx.MustGet(gin.BindKey).(*models.Discount)

		err := repo.CreateDiscount(discount)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func GetDiscount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		discount, err := repo.GetDiscount(map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(discount))
	}
}
func GetListDiscounts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		discounts, err := repo.GetListDiscounts()
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(discounts))
	}
}
func UpdateDiscount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		discount := ctx.MustGet(gin.BindKey).(*models.DiscountUpdate)

		err = repo.UpdateDiscount(id, discount)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
func DeleteDiscount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		err = repo.DeleteDiscount(id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
