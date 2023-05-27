package commandhandlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Passeder interface {
	GetPassedTime() string
}

func GetPassedHandler(p Passeder) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			p.GetPassedTime(),
		)
	}
}
