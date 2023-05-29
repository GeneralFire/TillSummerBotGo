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
	name    string
	timeNow time.Time
	want    summerTimeReturn
}

func TestGetSummerTime(t *testing.T) {
	domainInstance := timecalculator.New()

	testData := []testDataStruct{
		{
			name:    "First summer day",
			timeNow: time.Date(2000, time.June, 1, 0, 0, 0, 0, time.Local),
			want: summerTimeReturn{
				t:      timecalculator.TOTAL_SUMMER_TIME,
				summer: true,
			},
		},
		{
			name:    "One day before summer",
			timeNow: time.Date(2000, time.May, 31, 0, 0, 0, 0, time.Local),
			want: summerTimeReturn{
				t:      time.Hour * 24,
				summer: false,
			},
		},
		{
			name:    "One day after summer",
			timeNow: time.Date(2000, time.September, 1, 0, 0, 0, 0, time.Local),
			want: summerTimeReturn{
				t:      time.Hour*24*365 - timecalculator.TOTAL_SUMMER_TIME,
				summer: false,
			},
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			haveTime, haveSummer := domainInstance.GetSummerTime(tt.timeNow)
			assert.Equal(
				t,
				summerTimeReturn{
					t:      haveTime,
					summer: haveSummer,
				},
				tt.want,
			)
		})
	}
}
