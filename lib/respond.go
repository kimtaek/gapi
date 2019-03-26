package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"net/http"
)

type Data struct {
	Code          int
	Data          interface{}
	Message       string
	MessageCode   string
	PrefixMessage string
	SufFixMessage string
}

func Respond(c *gin.Context, data Data) {
	if data.Code == 0 {
		data.Code = 200
	}

	if data.Message == "" && data.MessageCode == "" {
		data.Message = http.StatusText(data.Code)
	}

	if data.Message == "" || data.MessageCode != "" {
		data.Message = getCustomMessage(data)
		data.Message = data.PrefixMessage + data.Message + data.SufFixMessage
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    data.Code,
		"message": data.Message,
		"data":    data.Data,
	})
	c.Abort()
	return
}

func getCustomMessage(data Data) string {
	if data.MessageCode != "" {
		return getCustomMessageWithCode(data.MessageCode)
	}
	var customMessage = ""
	switch data.Code {
	case 401:
		customMessage = "error message not found"
	}
	return customMessage
}

func getCustomMessageWithCode(code string) string {
	messageConfig, _ := ini.Load("messages/message.ini")
	return messageConfig.Section("ko").Key(code).String()
}
