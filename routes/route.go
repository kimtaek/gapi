package routes

import (
	"gapi/controllers"
	"gapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	systemCtl := new(controllers.System)
	authCtl := new(controllers.Auth)
	hotelCtl := new(controllers.Hotel)
	adminCtl := new(controllers.Admin)
	partnerAdminCtl := new(controllers.PartnerAdmin)
	codeCtl := new(controllers.Code)

	partnerRoute := r.Group("/ptn")
	{
		var authRoute *gin.RouterGroup

		partnerRoute.GET("systems/healthCheck", systemCtl.HealthCheck)
		// auth
		partnerRoute.POST("login", authCtl.Login)
		partnerRoute.Use(middlewares.AuthMiddleware())
		partnerRoute.POST("logout", authCtl.Logout)
		partnerRoute.GET("systems/config", systemCtl.GetConfig)

		authRoute = partnerRoute.Group("auth")
		{
			authRoute.GET("me", partnerAdminCtl.Me)
			authRoute.PUT("", partnerAdminCtl.Update)
			authRoute.PUT("password", partnerAdminCtl.UpdatePassword)
		}

		hotelRoute := partnerRoute.Group("hotels")
		{
			hotelRoute.GET("", hotelCtl.Index)
			hotelRoute.GET(":id", hotelCtl.Show)
		}

		codeRoute := partnerRoute.Group("codes")
		{
			codeRoute.GET(":code/options", codeCtl.Index)
		}
	}

	bmsRoute := r.Group("/bms")
	{
		bmsRoute.POST("login", authCtl.Login)
		bmsRoute.Use(middlewares.AuthMiddleware())
		bmsRoute.POST("logout", authCtl.Logout)

		authRoute := bmsRoute.Group("auth")
		{
			authRoute.GET("me", adminCtl.Me)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"Code": "404", "Message": "Page not found"})
	})

	return r
}
