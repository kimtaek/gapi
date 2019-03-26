package lib

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Slack struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
}

func SendSlackMessage(slack Slack) {
	slack.Username = strings.ToUpper(config.AppMode)
	if slack.Text == "" {
		return
	}
	if slack.Channel == "" {
		slack.Channel = "#exception"
	}
	params, _ := json.Marshal(slack)
	resp, _ := http.PostForm(
		config.AppConfig.SlackIncomingUrl,
		url.Values{
			"payload": {string(params)},
		},
	)
	defer resp.Body.Close()
}
