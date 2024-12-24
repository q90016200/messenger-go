package telgram

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Telegram struct {
	Token  string
	client *resty.Client
}

func NewTelegram(token string) *Telegram {
	return &Telegram{
		Token:  token,
		client: resty.New(),
	}
}

func (t *Telegram) SendMessage(channelID string, text string) error {
	requestUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.Token)

	client := t.client
	client = client.SetTimeout(time.Duration(10) * time.Second)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	params := map[string]string{
		"chat_id": channelID,
		"text":    text,
	}

	resp, err := client.R().
		SetHeaders(headers).
		SetFormData(params).Post(requestUrl)
	if err != nil {
		return err
	}
	return errors.New(fmt.Sprintf("\n\n[Telegram] sendMessage: %s\n\n", string(resp.Body())))
}