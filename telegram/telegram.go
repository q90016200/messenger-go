package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Telegram struct {
	botToken string
}

func NewTelegram(token string) *Telegram {
	return &Telegram{
		botToken: token,
	}
}

func (t *Telegram) Platform() string {
	return "Telegram"
}

func (t *Telegram) SendMessage(chatID string, text string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	payload := map[string]string{
		"chat_id":    chatID,
		"text":       text,
		"parse_mode": "MarkdownV2",
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to send message, response error: %s", err))
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("failed to send message, resp.StatusCode: %d ; resp: %s", resp.StatusCode, string(body)))
	}

	//fmt.Printf("\r\n[Telegram] sendMessage: %s \r\n", string(body))
	return nil
}
