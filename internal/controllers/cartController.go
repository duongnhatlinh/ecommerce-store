package controllers

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/middleware"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddItemToCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paramCart models.ParamCart
		err := ctx.BindQuery(&paramCart)
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		var data models.Cart
		data.UserId = user.Id
		data.ProductId = paramCart.ProductId
		data.Quantity = paramCart.Quantity

		if data.Quantity < 1 {
			data.Quantity = 1
		}

		err = repo.AddItemToCart(&data)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func RemoveItemFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		err = repo.RemoveItemFromCart(productId, user.Id)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func GetListItemsFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		cart, err := repo.GetListItemsFromCart(user.Id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(cart))
	}
}
