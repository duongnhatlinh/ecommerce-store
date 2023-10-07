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

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(gin.BindKey).(*models.User)

		err := repo.CreateUser(user)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSignIn := ctx.MustGet(gin.BindKey).(*models.UserSignIn)

		token, err := repo.AuthenticateUser(userSignIn)
		if err != nil || token == "" {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(token))
	}
}

func GetProfileUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(user))
	}
}

func UpdateInfoUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUpdate := ctx.MustGet(gin.BindKey).(*models.InfoUser)

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		err = repo.UpdateInfoUser(user.Id, userUpdate)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func UpdatePasswordUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		passwordUser := ctx.MustGet(gin.BindKey).(*models.UserPassword)

		user, err := middleware.CurrentUser(ctx)
		if err != nil {
			panic(err)
		}

		err = repo.UpdatePasswordUser(user, passwordUser)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		user, err := repo.GetUser(map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(user))
	}
}

func GetListUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging helper.Paging

		err := ctx.BindQuery(&paging)
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		users, err := repo.GetListUsers(&paging)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.NewSuccessRespond(users, paging, nil))
	}
}
