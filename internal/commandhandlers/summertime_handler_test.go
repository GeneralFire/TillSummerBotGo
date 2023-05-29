package commandhandlers

import (
	"fmt"
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

var (
	SAMPLE_DURATION = 3 * time.Hour
	SAMPLE_SUMMER   = false
)

type SummerTimeGetterStub struct{}

func (h *SummerTimeGetterStub) GetSummerTime(time.Time) (time.Duration, bool) {
	return SAMPLE_DURATION, SAMPLE_SUMMER
}

func TestSummertimeHandler(t *testing.T) {
	update := tgbotapi.Update{Message: &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}}
	passed := SummerTimeGetterStub{}
	handler := GetSummertimeHandler(&passed)
	msg := handler(update)

	assert.Equal(
		t,
		msg.Text,
		fmt.Sprintf("%v %v", SAMPLE_DURATION, SAMPLE_SUMMER),
	)
}
