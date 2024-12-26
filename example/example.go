package main

import (
	"fmt"
	"github.com/joho/godotenv"
	messenger "github.com/q90016200/messenger-go"
	"log"
	"os"
	"time"
)

func main() {
	manager := messenger.NewManager()

	// 先透過 .env 取得各平台所需參數
	// 加載 .env 檔案
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	testMsg := fmt.Sprintf("test message \r\n%d\r\n```%s```", time.Now().Unix(), "var a=\"test\"")

	// discord
	if os.Getenv("DISCORD_WEBHOOK_URL") != "" {
		discordMessenger := manager.Discord(os.Getenv("DISCORD_WEBHOOK_URL"))
		err := discordMessenger.SendMessage(testMsg)
		if err != nil {
			fmt.Println("[Discord] Error sending message:", err)
		}
	}

	// telegram
	if os.Getenv("TELEGRAM_BOT_TOKEN") != "" {
		err := manager.Telegram(os.Getenv("TELEGRAM_BOT_TOKEN")).SendMessage(os.Getenv("TELEGRAM_CHANNEL_ID"), testMsg)
		if err != nil {
			fmt.Println("[Telegram] Error sending message:", err)
		}
	}

	// line
	if os.Getenv("LINE_MESSAGE_ACCESS_TOKEN") != "" {
		err := manager.LineMessage(os.Getenv("LINE_MESSAGE_ACCESS_TOKEN")).TextMessage(os.Getenv("LINE_MESSAGE_CHANNEL_ID"), testMsg)
		if err != nil {
			fmt.Println("[Line-Message] Error sending message:", err)
		}
	}

}
