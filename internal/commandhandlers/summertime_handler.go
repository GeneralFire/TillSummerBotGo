package commandhandlers

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SummerTimeGetter interface {
	GetSummerTime(time.Time) string
}

func GetPassedHandler(p SummerTimeGetter) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			p.GetSummerTime(time.Now()),
		)
	}
}
