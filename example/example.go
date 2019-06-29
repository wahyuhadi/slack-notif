package main

import (
	send "github.com/wahyuhadi/slack-notif"
)


var (
	webhookUrl = "https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxx"
)


func main(){
	username := "bot_slack"
	channel := "general"
	msg := "Hello All , WelCome"

	// 1. webhookUrl
	// 2. msg
	// 3. channel
	// 4. username
	send.SendSlackNotification(webhookUrl, msg, channel,username)
}