package timecalculator

import (
	"time"
)

const (
	HOURS_IN_DAY = 24

	JUNE_DAYS_COUNT           = 30
	JULE_DAYS_COUNT           = 31
	AUGUST_DAYS_COUNT         = 31
	TOTAL_SUMMER_TIME_IN_DAYS = JUNE_DAYS_COUNT + JULE_DAYS_COUNT + AUGUST_DAYS_COUNT
	TOTAL_SUMMER_TIME         = TOTAL_SUMMER_TIME_IN_DAYS * HOURS_IN_DAY * time.Hour
)

type Domain struct{}

func New() Domain {
	return Domain{}
}

func (d *Domain) GetSummerTime(
	now time.Time,
) (time.Duration, bool) {
	if d.summer(now) {
		summerEndTime := time.Date(now.Year(), time.September, 1, 0, 0, 0, 0, now.Location())
		timePassed := summerEndTime.Sub(now)
		return timePassed, true
	}

	summerStartTime := time.Date(now.Year(), time.June, 1, 0, 0, 0, 0, now.Location())
	if now.After(summerStartTime) {
		summerStartTime = time.Date(now.Year()+1, time.June, 1, 0, 0, 0, 0, now.Location())
	}

	return summerStartTime.Sub(now), false
}

func (d *Domain) summer(
	now time.Time,
) bool {
	month := now.Month()
	if month >= time.June && month <= time.August {
		return true
	}
	return false
}

func (d *Domain) GetHello() string {
	return "Hello"
}
