package controllers

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/middleware"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userAddr := ctx.MustGet(gin.BindKey).(*models.UserAddress)

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}
		userAddr.UserId = user.Id

		err = repo.CreateUserAddress(userAddr)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func GetUserAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		userAddr, err := repo.GetUserAddress(map[string]interface{}{"user_id": user.Id})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(userAddr))
	}
}

func UpdateUserAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		userAddr := ctx.MustGet(gin.BindKey).(*models.UserAddressUpdate)

		err = repo.UpdateUserAddress(user.Id, userAddr)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
