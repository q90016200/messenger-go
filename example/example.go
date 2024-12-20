package main

import (
	"fmt"
	messenger "github.com/q90016200/messenger-go"
)

func main() {
	manager := messenger.NewManager()

	// 有效的 Discord Webhook URL
	discordMessenger, err := manager.Discord("https://discord.com/api/webhooks/1297750987032760430/2Cc0AdX_8wY0DYC2h9th_br_pum2fdpVtsKpmITghRtmwwvIpWZxnR5izVrC-bIkpFO2")
	if err != nil {
		fmt.Println("Error creating Discord messenger:", err)
		return
	}

	err = discordMessenger.SendMessage("test")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
