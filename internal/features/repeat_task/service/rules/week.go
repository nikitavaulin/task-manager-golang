package repeat_rules

import (
	"time"
)

type WeekRule struct {
	daysOfWeek map[time.Weekday]bool
}

func NewWeekRule(days map[time.Weekday]bool) (WeekRule, error) {
	return WeekRule{
		daysOfWeek: days,
	}, nil
}

func (r WeekRule) CalcNextDate(now time.Time, start time.Time) time.Time {
	if start.Before(now) {
		start = now
	}
	return r.closestDateFromDaysOfWeek(start)
}

func (r WeekRule) closestDateFromDaysOfWeek(start time.Time) time.Time {
	date := start
	for {
		date = date.AddDate(0, 0, 1)
		isIncluded := r.daysOfWeek[date.Weekday()]
		if isIncluded {
			return date
		}
	}
}
