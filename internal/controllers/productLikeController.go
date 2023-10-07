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

func CreateActionLike() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}
		productId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		var proLike *models.ProductLike
		proLike.UserId = user.Id
		proLike.ProductId = productId
		err = repo.CreateActionLike(proLike)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func CreateActionDislike() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}
		productId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		err = repo.CreateActionDislike(productId, user.Id)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
func GetListUsersLike() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		var paging helper.Paging

		err = ctx.BindQuery(&paging)
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		usersLike, err := repo.GetListUsersLike(&paging, productId)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.NewSuccessRespond(usersLike, paging, nil))
	}
}
