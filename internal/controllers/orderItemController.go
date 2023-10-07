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

func AddItemToOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		orderItem := ctx.MustGet(gin.BindKey).(*models.OrderItem)

		orderId, err := strconv.Atoi(ctx.Param("orderId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		orderItem.OrderId = orderId

		order, err := repo.AddItemToOrder(user.Id, orderItem)
		if err != nil {
			panic(err)
		}

		// update order
		totalPrice := order.TotalPrice + orderItem.Price*float32(orderItem.Quantity)

		err = repo.UpdateOrder(orderItem.OrderId, totalPrice)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func RemoveItemFromOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		orderId, err := strconv.Atoi(ctx.Param("orderId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		itemId, err := strconv.Atoi(ctx.Param("itemId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		// remove item
		err = repo.RemoveItemFromOrder(itemId, orderId, user.Id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
func GetListOrderItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		orderId, err := strconv.Atoi(ctx.Param("orderId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		orderItems, err := repo.GetListOrderItems(orderId)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(orderItems))
	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemId, err := strconv.Atoi(ctx.Param("itemId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		data := ctx.MustGet(gin.BindKey).(*models.OrderItemUpdate)

		err = repo.UpdateOrderItem(itemId, data)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))

	}
}
func GetOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemId, err := strconv.Atoi(ctx.Param("itemId"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		orderItem, err := repo.GetOrderItem(itemId)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(orderItem))
	}
}
