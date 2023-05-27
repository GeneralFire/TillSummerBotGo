package commandhandlers

import (
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

const PassedMsg = "Passed 3 days"

type SummerTimeGetterStub struct{}

func (h *SummerTimeGetterStub) GetSummerTime(time.Time) string {
	return PassedMsg
}
func TestPassedHandler(t *testing.T) {
	update := tgbotapi.Update{Message: &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}}
	passed := SummerTimeGetterStub{}
	handler := GetPassedHandler(&passed)
	msg := handler(update)

	assert.Equal(t, msg.Text, PassedMsg)
}
