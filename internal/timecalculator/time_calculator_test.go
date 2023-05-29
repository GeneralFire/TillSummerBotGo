package timecalculator_test

import (
	"testing"
	"time"

	"github.com/GeneralFire/TillSummerBotGo/internal/timecalculator"
	"github.com/stretchr/testify/assert"
)

type summerTimeReturn struct {
	t      time.Duration
	summer bool
}

type testDataStruct struct {
	timeNow time.Time
	want    summerTimeReturn
}

func TestGetSummerTime(t *testing.T) {
	domainInstance := timecalculator.New()

	testData := []testDataStruct{
		{
			timeNow: time.Date(2000, time.June, 1, 0, 0, 0, 0, time.Local),
			want: summerTimeReturn{
				t:      timecalculator.TOTAL_SUMMER_TIME,
				summer: true,
			},
		},
		{
			timeNow: time.Date(2000, time.May, 31, 0, 0, 0, 0, time.Local),
			want: summerTimeReturn{
				t:      time.Hour * 24,
				summer: false,
			},
		},
		{
			timeNow: time.Date(2000, time.September, 1, 0, 0, 0, 0, time.Local),
			want: summerTimeReturn{
				t:      time.Hour*24*365 - timecalculator.TOTAL_SUMMER_TIME,
				summer: false,
			},
		},
	}

	for _, data := range testData {
		// t.Parallel()
		haveTime, haveSummer := domainInstance.GetSummerTime(data.timeNow)
		assert.Equal(
			t,
			summerTimeReturn{
				t:      haveTime,
				summer: haveSummer,
			},
			data.want,
		)

	}
}
