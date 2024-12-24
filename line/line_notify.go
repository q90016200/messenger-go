package line

import (
	"fmt"
	"os/exec"
)

type LineNotify struct {
	Token string
}

func NewLineNotify(token string) *LineNotify {
	return &LineNotify{Token: token}
}

func (ln *LineNotify) SendMessage(message string) {
	cmd := exec.Command("curl", "-X", "POST", "-H", fmt.Sprintf("Authorization: Bearer %s", ln.Token), "-F", fmt.Sprintf("message=%s", message), "https://notify-api.line.me/api/notify")
	cmd.Start()
}
