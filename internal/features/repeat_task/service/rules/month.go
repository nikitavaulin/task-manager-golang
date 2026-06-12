package repeat_rules

import (
	"slices"
	"time"
)

type MonthRule struct {
	daysOfMonth             map[int]bool
	isLastDaySelected       bool
	isBeforeLastDaySelected bool
	months                  map[time.Month]bool
}

func NewMonthRule(
	daysOfMonth map[int]bool,
	isLastDaySelected bool,
	isBeforeLastDaySelected bool,
	months map[time.Month]bool,
) (MonthRule, error) {
	return MonthRule{
		daysOfMonth:             daysOfMonth,
		isLastDaySelected:       isLastDaySelected,
		isBeforeLastDaySelected: isBeforeLastDaySelected,
		months:                  months,
	}, nil
}

func (r MonthRule) CalcNextDate(now time.Time, start time.Time) time.Time {
	if start.Before(now) {
		start = now
	}
	return r.closestDayOfMonth(start)
}

func (r MonthRule) closestDayOfMonth(start time.Time) time.Time {
	date := start
	for {
		date = date.AddDate(0, 0, 1)
		if r.months == nil || r.months != nil && r.months[date.Month()] {
			if r.isDaySelected(date) {
				return date
			}
		}
	}
}

func (r MonthRule) isDaySelected(date time.Time) bool {
	day := date.Day()

	if r.daysOfMonth[day] {
		return true
	}

	if r.isLastDaySelected || r.isBeforeLastDaySelected {
		daysInMonth := countDaysInMonth(date.Month(), date.Year())
		diff := daysInMonth - day
		switch diff {
		case 0:
			return r.isLastDaySelected
		case 1:
			return r.isBeforeLastDaySelected
		}

	}

	return false
}

func countDaysInMonth(month time.Month, year int) int {
	m30 := []time.Month{time.April, time.June, time.September, time.November}

	switch {
	case month == time.February:
		if isLeapYear(year) {
			return 29
		}
		return 28

	case slices.Contains(m30, month):
		return 30

	default:
		return 31
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
