package commandhandlers

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

const HELLO_MSG = "Hallo"

type HelloerStub struct{}

func (h HelloerStub) GetHello() string {
	return HELLO_MSG
}
func TestHelloHandler(t *testing.T) {
	update := tgbotapi.Update{Message: &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}}
	halloer := HelloerStub{}

	handler := GetHelloHandler(halloer)
	msg := handler(update)

	assert.Equal(t, msg.Text, HELLO_MSG)
}
