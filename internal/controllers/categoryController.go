package controllers

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		category := ctx.MustGet(gin.BindKey).(*models.Category)

		err := repo.CreateCategory(category)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func GetCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}
		category, err := repo.GetCategory(map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(category))
	}
}
func GetListCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories, err := repo.GetListCategories()
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(categories))
	}
}
func UpdateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		category := ctx.MustGet(gin.BindKey).(*models.CategoryUpdate)

		err = repo.UpdateCategory(id, category)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
func DeleteCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		err = repo.DeleteCategory(id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
