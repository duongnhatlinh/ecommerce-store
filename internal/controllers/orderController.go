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

func CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order models.Order

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}
		order.UserId = user.Id
		order.TotalPrice = 0

		orderId, err := repo.CreateOrder(&order)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(orderId))
	}
}

func GetOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderId, err := strconv.Atoi(ctx.Param("orderId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		order, err := repo.GetOrder(orderId, user.Id)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(order))
	}
}

func GetListOrdersByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		orders, err := repo.GetListOrdersByUser(user.Id)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(orders))
	}
}

func GetTotalOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging helper.Paging

		err := ctx.BindQuery(&paging)
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		totalOrders, err := repo.GetTotalOrders(&paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.NewSuccessRespond(totalOrders, paging, nil))
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderId, err := strconv.Atoi(ctx.Param("orderId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		order := ctx.MustGet(gin.BindKey).(*models.Order)

		if order.TotalPrice <= 0 {
			panic(models.ErrOrderUpdateValueInvalid)
		}

		err = repo.UpdateOrder(orderId, order.TotalPrice)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderId, err := strconv.Atoi(ctx.Param("orderId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		err = repo.DeleteOrder(orderId, user.Id)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
