package controllers

import (
	"gapi/lib"
	"github.com/gin-gonic/gin"
)

type System struct {
}

func (_ *System) HealthCheck(c *gin.Context) {
	config := lib.InitConfigs()
	response := lib.Data{}
	response.Code = 200
	response.Message = "healthy"
	err := lib.Master().DB().Ping()
	if err != nil {
		response.Code = 400
		response.Message = err.Error()
		if config.AppMode == "release" {
			lib.SendSlackMessage(lib.Slack{
				Text: "HTTP HEALTH CHECK(DB)::" + err.Error(),
			})
		}
		c.JSON(response.Code, response)
		c.Abort()
		return
	}
	err = lib.RedisPing()
	if err != nil {
		response.Code = 400
		response.Message = err.Error()
		if config.AppMode == "release" {
			lib.SendSlackMessage(lib.Slack{
				Text: "HTTP HEALTH CHECK(REDIS)::" + err.Error(),
			})
		}
		c.JSON(response.Code, response)
		c.Abort()
		return
	}
	lib.Respond(c, response)
}

func (_ *System) GetConfig(c *gin.Context) {
	response := lib.Data{}
	response.Data = lib.GetConfigs()
	lib.Respond(c, response)
}
