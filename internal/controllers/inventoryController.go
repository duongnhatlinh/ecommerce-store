package controllers

import (
	"ecommercestore/internal/helper"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		inventory := ctx.MustGet(gin.BindKey).(*models.Inventory)

		err := repo.CreateInventory(inventory)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}

func GetInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		inventory, err := repo.GetInventory(map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(inventory))
	}
}
func GetListInventories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		inventories, err := repo.GetListInventories()
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(inventories))
	}
}
func UpdateInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		inventory := ctx.MustGet(gin.BindKey).(*models.Inventory)

		err = repo.UpdateInventory(id, inventory)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
func DeleteInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(helper.ErrInvalidRequest(err))
		}

		err = repo.DeleteInventory(id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, helper.SimpleSuccessRespond(true))
	}
}
