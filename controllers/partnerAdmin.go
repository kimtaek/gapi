package controllers

import (
	"gapi/lib"
	"gapi/models"
	"github.com/gin-gonic/gin"
	"strings"
)

type PartnerAdmin struct {
	models.PartnerAdmin
}

type ChangePassword struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func (_ *PartnerAdmin) Me(c *gin.Context) {
	var response = lib.Data{}
	authorization := c.GetHeader("Authorization")
	token := strings.Split(authorization, "Bearer ")[1]
	claims, err := lib.ParseToken(token, c)
	if err == nil {
		model := new(models.PartnerAdmin)
		user := model.Find(claims.Auth.Id)
		response.Data = user
	}
	lib.Respond(c, response)
}

func (_ *PartnerAdmin) Update(c *gin.Context) {
	response, request, auth := lib.Data{}, models.PartnerAdmin{}, lib.AuthUserInfo(c)
	c.BindJSON(&request)

	model := new(models.PartnerAdmin)
	model.Update(auth.Id, request)
	lib.Respond(c, response)
}

func (_ *PartnerAdmin) UpdatePassword(c *gin.Context) {
	response, request, auth := lib.Data{}, ChangePassword{}, lib.AuthUserInfo(c)
	c.BindJSON(&request)

	db := lib.Master()
	user := models.PartnerAdmin{}
	db.Model(models.PartnerAdmin{}).Find(&user, "id = ?", auth.Id)
	if user.Id == 0 {
		response.Code = 400
		response.MessageCode = "request.not.complete"
		lib.Respond(c, response)
		return
	}

	if lib.ComparePassword(user.Password, request.Old) == false {
		response.Code = 400
		response.MessageCode = "auth.mismatch.password"
		lib.Respond(c, response)
		return
	}

	user.Password = lib.GeneratePassword(request.New)
	db.Model(models.PartnerAdmin{}).Where("id = ?", auth.Id).Update(&user)
	lib.Respond(c, response)
}
