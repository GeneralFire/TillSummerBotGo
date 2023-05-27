package commandhandlers

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

const PassedMsg = "Passed 3 days"

type PassedStub struct{}

func (h *PassedStub) GetPassedTime() string {
	return PassedMsg
}
func TestPassedHandler(t *testing.T) {
	update := tgbotapi.Update{Message: &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}}
	passed := PassedStub{}
	handler := GetPassedHandler(&passed)
	msg := handler(update)

	assert.Equal(t, msg.Text, PassedMsg)
}
