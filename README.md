# Simple Slack WebHook

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

### example import package(s)
```go
package main

import (
	send "github.com/wahyuhadi/slack-notif"
)

```


### example main
```go
var (
	// URL webhook 
	webhookUrl = "https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxx"
)

func main(){
	username := "bot_slack"
	channel := "general"
	msg := "Hello All , WelCome"
	send.SendSlackNotification(webhookUrl, msg, channel,username)
}
```

### important 
```go
// 1. webhookUrl
// 2. msg
// 3. channel
// 4. username
send.SendSlackNotification(webhookUrl, msg, channel,username)
```


**Free Software, Yeah!**