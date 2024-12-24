package main

import (
	"fmt"
	messenger "github.com/q90016200/messenger-go"
)

func main() {
	manager := messenger.NewManager()

	// 有效的 Discord Webhook URL
	discordMessenger := manager.Discord("webhookUrl")
	err := discordMessenger.SendMessage("test")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
