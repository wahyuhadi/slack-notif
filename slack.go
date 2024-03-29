package service


import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "time"
)

type SlackRequestBody struct {
	Text string `json:"text"`
	Channel  string `json:"channel"`
	Username  string `json:"username"`
}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func SendSlackNotification(webhookUrl, msg, channel, username string) error {
    slackBody, _ := json.Marshal(SlackRequestBody{Text: msg, Channel:channel, Username:username})
    req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
    if err != nil {
        return err
    }

    req.Header.Add("Content-Type", "application/json")

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    if buf.String() != "ok" {
        return errors.New("Non-ok response returned from Slack")
    }
    return nil
}