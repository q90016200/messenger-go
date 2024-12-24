package messenger_go

import (
	"github.com/q90016200/messenger-go/discord"
	"github.com/q90016200/messenger-go/line"
)

type Messenger interface {
	SendMessage(channelID, message string) error
	Platform() string // 返回平台名稱
}

type Manager struct {
	//line     Messenger
	lineMessage Messenger
	lineNotify  Messenger
	telegram    Messenger
	discord     *discord.Discord
}

// NewManager 初始化 Manager
func NewManager() *Manager {
	return &Manager{}
}

// 各平台具體方法
//func (m *Manager) Line() Messenger {
//	return m.line
//}

//func (m *Manager) LineMessage() Messenger {
//	return m.line
//}

func (m *Manager) LineNotify(token string) *line.LineNotify {
	return line.NewLineNotify(token)
}

func (m *Manager) Discord(webhookUrl string) *discord.Discord {
	return discord.NewDiscord(webhookUrl)
}

func (m *Manager) Telegram() Messenger {
	return m.telegram
}
