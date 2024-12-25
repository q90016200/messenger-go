package line

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type LineMessage struct {
	channelAccessToken string
	request            *resty.Request
}

func NewLineMessage(channelAccessToken string) *LineMessage {
	client := resty.New()
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", channelAccessToken)
	request := client.R().SetHeaders(headers)

	return &LineMessage{
		channelAccessToken: channelAccessToken,
		request:            request,
	}
}

type textMessage struct {
	To       string `json:"to"`
	Messages []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
}

func (lm *LineMessage) TextMessage(channelID string, message string) error {
	payload := textMessage{
		To: channelID,
		Messages: []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}{
			{Type: "text", Text: message},
		},
	}

	resp, err := lm.request.SetBody(payload).
		Post("https://api.line.me/v2/bot/message/push")

	if err != nil {
		return errors.New(fmt.Sprintf("\n\n[Line-Message] TextMessage err: %s\n\n", err.Error()))
	}
	fmt.Printf("\n\n[Line-Message] sendMessage: %s\n\n", string(resp.Body()))
	return nil
}
