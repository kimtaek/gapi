package controllers

import (
	"errors"
	"gapi/lib"
	"gapi/models"
	"github.com/gin-gonic/gin"
	"strings"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (_ *Auth) Login(c *gin.Context) {
	response, request := lib.Data{}, Auth{}
	c.BindJSON(&request)

	issuer := strings.Split(c.Request.URL.Path, "/")[1]

	var err error
	var auth lib.JwtAuth
	switch issuer {
	case "ptn":
		auth, err = partnerLogin(request.Email, request.Password)
	case "bms":
		auth, err = adminLogin(request.Email, request.Password)
	default:
		response.Code = 500
		lib.Respond(c, response)
		return
	}

	if err != nil {
		response.Code = 400
		response.MessageCode = err.Error()
		lib.Respond(c, response)
		return
	}

	token, expiredAt, err := lib.GenerateToken(auth, c)
	if err != nil {
		response.Code = 400
		response.MessageCode = "auth.token.not.generate"
		lib.Respond(c, response)
		return
	}

	data := make(map[string]interface{})
	data["Token"] = token
	data["TokenType"] = "Bearer"
	data["ExpiredAt"] = expiredAt

	response.Data = data
	lib.Respond(c, response)
}

func (_ *Auth) Logout(c *gin.Context) {
	var response = lib.Data{}
	authorization := c.GetHeader("Authorization")
	token := strings.Split(authorization, "Bearer ")[1]
	claims, err := lib.ParseToken(token, c)
	if err == nil {
		lib.DeleteRedis(claims.Issuer + ":" + token)
	}
	lib.Respond(c, response)
}

func adminLogin(email string, password string) (lib.JwtAuth, error) {
	model := new(models.Admin)
	var jwtAuth lib.JwtAuth
	user := model.FindByEmail(email)
	if user.Id == 0 {
		return jwtAuth, errors.New("auth.mismatch.id")
	}

	isPassed := lib.ComparePassword(user.Password, password)
	if isPassed == false {
		return jwtAuth, errors.New("auth.mismatch.password")
	}

	jwtAuth.Id = user.Id
	jwtAuth.Email = user.Email
	jwtAuth.Name = user.Name
	jwtAuth.Role = "admin"
	jwtAuth.Issuer = "bms"
	return jwtAuth, nil
}

func partnerLogin(email string, password string) (lib.JwtAuth, error) {
	var jwtAuth lib.JwtAuth
	var partnerAdminModel = new(models.PartnerAdmin)
	var partnerAdminServiceModel = new(models.PartnerAdminService)

	user := partnerAdminModel.FindByEmail(email)
	if user.Id == 0 {
		return jwtAuth, errors.New("auth.mismatch.id")
	}

	isPassed := lib.ComparePassword(user.Password, password)
	if isPassed == false {
		return jwtAuth, errors.New("auth.mismatch.password")
	}

	var hotels []int
	where := models.PartnerAdminService{
		PartnerAdminId: user.Id,
	}

	services := partnerAdminServiceModel.WhereIndex(where)
	for _, service := range services {
		if service.Service == "HOTEL" {
			hotels = append(hotels, service.ServiceId)
		}
	}

	jwtAuth.Id = user.Id
	jwtAuth.Email = user.Email
	jwtAuth.Name = *user.Name
	jwtAuth.Role = "master"
	jwtAuth.Issuer = "ptn"
	jwtAuth.Service.Hotels = hotels
	return jwtAuth, nil
}
