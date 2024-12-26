package line

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LineMessage struct {
	channelAccessToken string
}

func NewLineMessage(channelAccessToken string) *LineMessage {
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", channelAccessToken)

	return &LineMessage{
		channelAccessToken: channelAccessToken,
	}
}

type textMessage struct {
	To       string `json:"to"`
	Messages []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
}

func (lm *LineMessage) Platform() string {
	return "LineMessage"
}

func (lm *LineMessage) TextMessage(channelID string, message string) error {
	url := "https://api.line.me/v2/bot/message/push"
	payload := textMessage{
		To: channelID,
		Messages: []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}{
			{Type: "text", Text: message},
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", lm.channelAccessToken))

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
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
