package commandhandlers_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/GeneralFire/TillSummerBotGo/internal/commandhandlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestSummertimeHandler(t *testing.T) {
	update := tgbotapi.Update{Message: &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}}

	type mockReturnStruct struct {
		t time.Duration
		b bool
	}
	type handlerOutputStruct struct {
		msg string
	}

	mockReturn := []mockReturnStruct{
		{
			t: time.Hour*25 + time.Minute*5 + time.Second*6, // 1 day, 25:05:06
			b: false,
		},
		{
			t: time.Hour*5 + time.Minute*25 + time.Second*25, // 0 day, 05:25:25
			b: false,
		},
		{
			t: time.Hour*19 + time.Minute*59 + time.Second*25,
			b: true,
		},
		{
			t: time.Hour*49 + time.Minute*15 + time.Second*13,
			b: true,
		},
		{
			t: time.Hour*2207 + time.Minute*15 + time.Second*13,
			b: true,
		},
		{
			t: time.Hour*0 + time.Minute*1 + time.Second*59, // 91, 0.98
			b: true,
		},
	}

	want := []handlerOutputStruct{
		{
			msg: fmt.Sprintf(commandhandlers.UNTIL_SUMMER, 1, "25:05:06"),
		},
		{
			msg: fmt.Sprintf(commandhandlers.UNTIL_SUMMER, 0, "05:25:25"),
		},
		{
			msg: fmt.Sprintf(commandhandlers.UNTIL_SUMMER_END, 0, 99.09),
		},
		{
			msg: fmt.Sprintf(commandhandlers.UNTIL_SUMMER_END, 2, 97.77),
		},
		{
			msg: fmt.Sprintf(commandhandlers.UNTIL_SUMMER_END, 91, 0.03),
		},
		{
			msg: fmt.Sprintf(commandhandlers.UNTIL_SUMMER_END, 0, 100.00),
		},
	}

	assert.Equal(t, len(want), len(mockReturn))

	for i := 0; i < len(mockReturn); i++ {
		summerTimeGetter := commandhandlers.NewSummerTimeGetterMock(t)
		summerTimeGetter.GetSummerTimeMock.Return(
			mockReturn[i].t, mockReturn[i].b,
		)
		handler := commandhandlers.GetSummertimeHandler(summerTimeGetter)

		have := handler(update)
		assert.Equal(t, have.Text, want[i].msg)
	}

}
