package main

import (
	"fmt"
	"github.com/joho/godotenv"
	messenger "github.com/q90016200/messenger-go"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	manager := messenger.NewManager()

	// 先透過 .env 取得各平台所需參數
	_, filePath, _, _ := runtime.Caller(0)                     // 獲取當前文件的絕對路徑
	projectRoot := filepath.Join(filepath.Dir(filePath), "..") // 假設 .env 位於項目根目錄
	envPath := filepath.Join(projectRoot, ".env")
	// 加載 .env 檔案
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	testMsg := fmt.Sprintf("test message \r\n%d\r\n```%s```", time.Now().Unix(), "var a=\"test\"")

	// discord
	if os.Getenv("DISCORD_WEBHOOK_URL") != "" {
		discordMessenger := manager.Discord(os.Getenv("DISCORD_WEBHOOK_URL"))
		err := discordMessenger.SendMessage(testMsg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

	// telegram
	if os.Getenv("TELEGRAM_BOT_TOKEN") != "" {
		err := manager.Telegram(os.Getenv("TELEGRAM_BOT_TOKEN")).SendMessage(os.Getenv("TELEGRAM_CHANNEL_ID"), testMsg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

	// line
	if os.Getenv("LINE_MESSAGE_ACCESS_TOKEN") != "" {
		err := manager.LineMessage(os.Getenv("LINE_MESSAGE_ACCESS_TOKEN")).TextMessage(os.Getenv("LINE_MESSAGE_CHANNEL_ID"), testMsg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

}
