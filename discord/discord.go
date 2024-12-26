package discord

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"time"
)

type Discord struct {
	webHookUrl string
	client     *resty.Client
}

func NewDiscord(webHookUrl string) *Discord {
	client := resty.New()

	return &Discord{
		webHookUrl: webHookUrl,
		client:     client,
	}
}

func (d *Discord) SendMessage(message string) error {
	err := validateWebhookURL(d.webHookUrl)
	if err != nil {
		return err
	}

	payload := map[string]string{"content": message}
	client := d.client
	client = client.SetTimeout(time.Duration(10) * time.Second)
	resp, err := client.R().
		SetBody(payload).
		Post(d.webHookUrl)
	if err != nil {
		return errors.New(fmt.Sprintf("[Discord] failed to send message err: %v", err))
	}
	if resp.StatusCode() >= 400 {
		return errors.New(fmt.Sprintf("[Discord] failed to send message status: %s", resp.Status()))
	}

	fmt.Printf("\n\n[Discord] sendMessage: %s\n\n", string(resp.Body()))
	return nil
}

//func (d *Discord) SendEmbed() {
//	err := validateWebhookURL(d.WebHookUrl)
//	if err != nil {
//		return err
//	}
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
