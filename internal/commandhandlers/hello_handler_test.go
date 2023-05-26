package commandhandlers

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	update := tgbotapi.Update{Message: &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}}
	msg := HelloHandler(update)

	assert.Equal(t, msg.Text, "Hello")
}
