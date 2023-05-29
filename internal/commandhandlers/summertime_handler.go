package commandhandlers

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SummerTimeGetter interface {
	GetSummerTime(time.Time) (time.Duration, bool)
}

func GetPassedHandler(p SummerTimeGetter) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		t, b := p.GetSummerTime(time.Now())
		msgText := fmt.Sprintf("%v %v", t, b)
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			msgText,
		)
	}
}
