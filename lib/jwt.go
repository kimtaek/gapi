package lib

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

type Service struct {
	Hotels []int
}

type JwtAuth struct {
	Id      int
	Role    string
	Service Service
	Email   string
	Name    string
	Issuer  string
}

type Claims struct {
	Auth JwtAuth
	jwt.StandardClaims
}

func GenerateToken(auth JwtAuth, ctx *gin.Context) (string, int64, error) {
	authorization := ctx.GetHeader("Authorization")

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.AppConfig.JwtTTL) * time.Hour)
	claims := Claims{
		auth,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    prefix(ctx.Request.URL.Path, 1),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.AppConfig.AppSecret))
	if err != nil {
		return "", 0, err
	}

	if authorization != "" {
		RefreshToken(auth, claims.Issuer, token, ctx)
		log.Println("Token Refreshed With Redis")
	}

	duration := (time.Duration(config.AppConfig.JwtTTL) * time.Hour).Seconds()
	SetCache(claims.Issuer+":"+token, []byte(""), duration)
	return token, expireTime.Unix(), nil
}

func RefreshToken(_ JwtAuth, Issuer string, token string, ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		//return nil
		return
	}
	requestToken := strings.Split(authorization, "Bearer ")[1]
	// todo:: set token expired at now time +3 minutes and save new token
	duration := (time.Duration(30) * time.Second).Seconds()
	SetCache(Issuer+":"+requestToken, []byte(""), duration)
}

func ParseToken(token string, ctx *gin.Context) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.AppSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			tokenData, err := GetCache(claims.Issuer + ":" + token)
			if err == nil && tokenData != nil {
				return claims, nil
			}
			return nil, err
		}
	}

	return nil, err
}

func prefix(segments string, index int) string {
	return strings.Split(segments, "/")[index]
}

func AuthUserInfo(c *gin.Context) JwtAuth {
	var auth JwtAuth
	authorization := c.GetHeader("Authorization")
	token := strings.Split(authorization, "Bearer ")[1]
	claims, err := ParseToken(token, c)
	if err == nil {
		auth = claims.Auth
	}
	return auth
}
