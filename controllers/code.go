package controllers

import (
	"gapi/lib"
	"gapi/models"
	"github.com/gin-gonic/gin"
)

type Code struct {
}

func (_ *Code) Index(c *gin.Context) {
	var response = lib.Data{}
	groupCode := c.Param("code")
	model := new(models.Code)

	where := models.Code{
		Group: groupCode,
	}
	response.Data = model.WhereFindOptions(where)
	lib.Respond(c, response)
}
