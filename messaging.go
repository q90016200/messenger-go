package messenger_go

import "github.com/q90016200/messenger-go/discord"

type Messenger interface {
	SendMessage(channelID, message string) error
	Platform() string // 返回平台名稱
}

type Manager struct {
	line     Messenger
	telegram Messenger
	discord  *discord.Discord
}

// NewManager 初始化 Manager
func NewManager() *Manager {
	return &Manager{}
}

// 各平台具體方法
func (m *Manager) Line() Messenger {
	return m.line
}

func (m *Manager) Discord(webhookUrl string) (*discord.Discord, error) {
	return discord.NewDiscord(webhookUrl)
}

func (m *Manager) Telegram() Messenger {
	return m.telegram
}
