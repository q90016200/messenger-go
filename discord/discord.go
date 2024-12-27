package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Discord struct {
	webHookUrl string
}

func NewDiscord(webHookUrl string) *Discord {

	return &Discord{
		webHookUrl: webHookUrl,
	}
}

func (d *Discord) Platform() string {
	return "Discord"
}

func (d *Discord) SendMessage(message string) error {
	err := validateWebhookURL(d.webHookUrl)
	if err != nil {
		return err
	}

	payload := map[string]string{"content": message}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", d.webHookUrl, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Discord response error: %s", err))
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New(fmt.Sprintf("failed to send message, resp.StatusCode: %d ; resp: %s", resp.StatusCode, string(body)))
	}

	fmt.Printf("\r\n[Discord] sendMessage successfully \r\n")
	return nil
}

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
