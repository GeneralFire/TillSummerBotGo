package commandhandlers

//go:generate minimock -i SummerTimeGetter -o ./ -s "_minimock.go"

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SummerTimeGetter interface {
	GetSummerTime(time.Time) (time.Duration, bool)
}

const (
	UNTIL_SUMMER     = "%d full days left until summer (or %s)"
	UNTIL_SUMMER_END = "%d full days left until the end of summer(%3.2f%% complete)"

	HOURS_IN_DAY = 24

	JUNE_DAYS_COUNT           = 30
	JULE_DAYS_COUNT           = 31
	AUGUST_DAYS_COUNT         = 31
	TOTAL_SUMMER_TIME_IN_DAYS = JUNE_DAYS_COUNT + JULE_DAYS_COUNT + AUGUST_DAYS_COUNT
	TOTAL_SUMMER_TIME         = TOTAL_SUMMER_TIME_IN_DAYS * HOURS_IN_DAY * time.Hour
)

func GetSummertimeHandler(p SummerTimeGetter) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		timeLeft, summer := p.GetSummerTime(time.Now().Add(time.Hour * 3))
		days := timeLeft / (time.Hour * HOURS_IN_DAY)
		var msgText string

		if summer {
			msgText = fmt.Sprintf(
				UNTIL_SUMMER_END,
				days,
				(1-timeLeft.Hours()/TOTAL_SUMMER_TIME.Hours())*100,
			)
		} else {
			var hoursAsString, minutesAsString, secondsAsString string
			hours := int32(timeLeft.Hours())
			if hours < 10 {
				hoursAsString = fmt.Sprintf("0%d", hours)
			} else {
				hoursAsString = fmt.Sprintf("%d", hours)
			}
			minutes := int32(timeLeft.Minutes()) % 60
			if minutes < 10 {
				minutesAsString = fmt.Sprintf("0%d", minutes)
			} else {
				minutesAsString = fmt.Sprintf("%d", minutes)
			}
			seconds := int32(timeLeft.Seconds()) % 60
			if seconds < 10 {
				secondsAsString = fmt.Sprintf("0%d", seconds)
			} else {
				secondsAsString = fmt.Sprintf("%d", seconds)
			}

			msgText = fmt.Sprintf(
				UNTIL_SUMMER,
				days,
				fmt.Sprintf("%s:%s:%s", hoursAsString, minutesAsString, secondsAsString), // timeLeft.Hours(),
			)
		}
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			msgText,
		)
	}
}
