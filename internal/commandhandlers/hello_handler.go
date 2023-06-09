package commandhandlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Helloer interface {
	GetHello() string
}

func GetHelloHandler(halloer Helloer) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			halloer.GetHello(),
		)
	}
}
