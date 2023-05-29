package commandhandlers

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SummerTimeGetter interface {
	GetSummerTime(time.Time) (time.Duration, bool)
}

const (
	UNTIL_SUMMER     = "%d days left until summer (or %.2f hours)"
	UNTIL_SUMMER_END = "%d dyas left until the end of summer(%.2f)"

	HOURS_IN_DAY = 24

	JUNE_DAYS_COUNT           = 30
	JULE_DAYS_COUNT           = 31
	AUGUST_DAYS_COUNT         = 31
	TOTAL_SUMMER_TIME_IN_DAYS = JUNE_DAYS_COUNT + JULE_DAYS_COUNT + AUGUST_DAYS_COUNT
	TOTAL_SUMMER_TIME         = TOTAL_SUMMER_TIME_IN_DAYS * HOURS_IN_DAY * time.Hour
)

func GetSummertimeHandler(p SummerTimeGetter) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		timeLeft, summer := p.GetSummerTime(time.Now())
		days := timeLeft / (time.Hour * HOURS_IN_DAY)
		var msgText string

		if summer {
			msgText = fmt.Sprintf(
				UNTIL_SUMMER_END,
				days,
				1-timeLeft.Hours()/TOTAL_SUMMER_TIME.Hours(),
			)
		} else {
			msgText = fmt.Sprintf(UNTIL_SUMMER, days, timeLeft.Hours())
		}
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			msgText,
		)
	}
}
