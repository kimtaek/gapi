package main

import (
	"gapi/lib"
	"gapi/routes"
	"github.com/gin-gonic/gin"
	"io"
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
	f, _ := os.Create(config.Logs.System)
	gin.DefaultWriter = io.MultiWriter(f)
	routes.RegisterRoutes().Run()
}
