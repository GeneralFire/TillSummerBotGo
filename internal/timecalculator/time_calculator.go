package timecalculator

import (
	"time"
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
