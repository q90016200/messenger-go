package telegram

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func (t *Telegram) SendMessage(chatId string, text string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	payload := url.Values{
		"chat_id":    {chatId},
		"text":       {text},
		"parse_mode": {"MarkdownV2"},
	}

	resp, err := client.PostForm(apiURL, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("failed to send message, resp.StatusCode: %d", resp.StatusCode))
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to send message, response error: %s", err))
	}

	fmt.Printf("\r\n[Telegram] sendMessage: %s\r\n", string(body))
	return nil
}
