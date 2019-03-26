package main

import (
	"gapi/lib"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	config := lib.InitConfigs()
	db := lib.InitDatabase()
	redis := lib.InitRedis()
	defer db.Close()
	defer redis.Close()

	if config.AppMode == "" {
		log.Println("APP_MODE:: " + config.AppMode)
		os.Exit(1)
	}
	gin.SetMode(config.AppMode)
	err := lib.Master().DB().Ping()
	if err != nil {
		if config.AppMode == "release" {
			lib.SendSlackMessage(lib.Slack{
				Text: "DOCKER HEALTH CHECK(DB)::" + err.Error(),
			})
		}
		log.Println("DOCKER HEALTH CHECK(DB)::" + err.Error())
		os.Exit(1)
	}
	err = lib.RedisPing()
	if err != nil {
		if config.AppMode == "release" {
			lib.SendSlackMessage(lib.Slack{
				Text: "DOCKER HEALTH CHECK(REDIS)::" + err.Error(),
			})
		}
		log.Println("DOCKER HEALTH CHECK(REDIS)::" + err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
