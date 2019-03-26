package middlewares

import (
	"gapi/lib"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = lib.Data{}
		response.Code = 401
		authorization := c.GetHeader("Authorization")
		if authorization == "" || strings.Contains(authorization, "Bearer ") == false {
			response.MessageCode = "auth.token.not.require"
			lib.Respond(c, response)
			return
		}

		token := strings.Split(authorization, "Bearer ")[1]
		if token == "" {
			response.MessageCode = "auth.token.not.require"
			lib.Respond(c, response)
			return
		}
		claims, err := lib.ParseToken(token, c)
		if err != nil {
			response.MessageCode = "auth.token.not.parsed"
			lib.Respond(c, response)
			return
		}
		firstUrlSegment := strings.Split(c.Request.URL.Path, "/")[1]
		if claims.Issuer != firstUrlSegment {
			response.MessageCode = "auth.token.issuer.mismatched"
			lib.Respond(c, response)
			return
		}

		nowTime := time.Now()
		config := lib.GetConfigs()
		refreshTime := time.Unix(claims.ExpiresAt, 0).Add(-time.Duration(config.AppConfig.JWtTRT) * time.Minute)

		log.Println("Token Refresh At: " + refreshTime.String())

		if nowTime.Unix() > refreshTime.Unix() {
			token, expiredAt, err := lib.GenerateToken(claims.Auth, c)
			if err == nil {
				c.Header("Authorization", token)
				log.Println("NowTime: " + nowTime.String())
				log.Println("Token Refreshed: ")
				log.Println("New Token: " + token)
				log.Println("New Token Expired Time: " + time.Unix(expiredAt, 0).String())
			} else {
				lib.Respond(c, response)
				return
			}
		}
		c.Next()
	}
}
