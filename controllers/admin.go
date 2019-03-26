package controllers

import (
	"gapi/lib"
	"gapi/models"
	"github.com/gin-gonic/gin"
	"strings"
)

type Admin struct {
	models.Admin
}

func (_ *Admin) Me(c *gin.Context) {
	var response = lib.Data{}
	authorization := c.GetHeader("Authorization")
	token := strings.Split(authorization, "Bearer ")[1]
	claims, err := lib.ParseToken(token, c)
	if err == nil {
		model := new(models.Admin)
		user := model.Find(claims.Auth.Id)
		response.Data = user
	}
	lib.Respond(c, response)
}
