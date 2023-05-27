package domain

import (
	"fmt"
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
	// format string,
) string {
	if d.summer(now) {
		// summerStartTime := time.Date(now.Year(), time.June, 1, 0, 0, 0, 0, now.Location())
		summerEndTime := time.Date(now.Year(), time.September, 1, 0, 0, 0, 0, now.Location())
		timePassed := summerEndTime.Sub(now)
		passedPercent := float32(timePassed) / float32(TOTAL_SUMMER_TIME) * 100
		// timePassed := (now.Add(-time.Duration(now.Year() * time.Duration.)))
		dayCount := int(timePassed / (time.Hour * HOURS_IN_DAY))
		return fmt.Sprintf(
			"Summer passed %d days or %.2f %%", dayCount, passedPercent,
		)
	}
	return d.getTimeTillSummer(now)
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

func (d *Domain) getTimeTillSummer(currentTime time.Time) string {
	return "3 days"
}

func (d *Domain) GetHello() string {
	return "Hello"
}

func (d *Domain) GetPassedTime() string {
	return "3 days"
}
