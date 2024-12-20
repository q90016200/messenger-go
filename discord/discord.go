package discord

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
)

type Discord struct {
	WebHookUrl string
	client     *resty.Client
}

func NewDiscord(webHookUrl string) (*Discord, error) {
	err := validateWebhookURL(webHookUrl)
	if err != nil {
		return nil, err
	}

	client := resty.New()

	return &Discord{
		WebHookUrl: webHookUrl,
		client:     client,
	}, nil
}

func (d *Discord) SendMessage(message string) error {
	payload := map[string]string{"content": message}
	resp, err := d.client.R().
		SetBody(payload).
		Post(d.WebHookUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 400 {
		return fmt.Errorf("failed to send message: %s", resp.Status())
	}

	return nil
}

//func (d *Discord) SendEmbed() {
//
//}

// 验证 webhook URL 是否有效
func validateWebhookURL(webhookURL string) error {
	if webhookURL == "" {
		return errors.New("webhook URL cannot be empty")
	}
	if _, err := url.ParseRequestURI(webhookURL); err != nil {
		return fmt.Errorf("invalid webhook URL: %w", err)
	}
	return nil
}
