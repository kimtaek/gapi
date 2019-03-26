package controllers

import (
	"gapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Hotel struct {
}

func (_ *Hotel) Index(ctx *gin.Context) {
	model := new(models.Hotel)
	data := model.Index()
	ctx.JSON(http.StatusOK, gin.H{
		"hotels": data,
	})
}

func (_ *Hotel) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	model := new(models.Hotel)
	data := model.Find(id)
	ctx.JSON(http.StatusOK, gin.H{
		"hotels": data,
	})
}

func (_ *Hotel) ShowBrandHotel(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	model := new(models.Hotel)
	data := model.WhereFind(models.Hotel{Brand: id})
	ctx.JSON(http.StatusOK, gin.H{
		"hotels": data,
	})
}

func (_ *Hotel) IndexBrandHotel(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	model := new(models.Hotel)
	data := model.WhereIndex(models.Hotel{Brand: id})
	ctx.JSON(http.StatusOK, gin.H{
		"hotels": data,
	})
}

func (_ *Hotel) IndexWhere(ctx *gin.Context) {
	model := new(models.Hotel)
	data := model.WhereIndex(models.Hotel{
		Provider: 1,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"hotels": data,
	})
}
