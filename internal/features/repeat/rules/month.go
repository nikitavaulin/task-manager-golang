package repeat_rules

import (
	"slices"
	"time"
)

type MonthRule struct {
	daysOfMonth map[int]bool
	months      map[time.Month]bool
}

func NewMonthRule(
	daysOfMonth map[int]bool,
	months map[time.Month]bool,
) (MonthRule, error) {
	return MonthRule{
		daysOfMonth: daysOfMonth,
		months:      months,
	}, nil
}

func (r MonthRule) CalcNextDate(now time.Time, start time.Time) time.Time {
	if start.Before(now) {
		start = now
	}
	return r.closestDayOfMonth(start)
}

func (r MonthRule) closestDayOfMonth(start time.Time) time.Time {
	if r.months == nil {
		return r.closestDayOfAnyMonth(start)
	}
	return r.closestDayOfCertainMonths(start)
}

func (r MonthRule) closestDayOfAnyMonth(start time.Time) time.Time {
	date := start
	for {
		if r.daysOfMonth[date.Day()] {
			return date
		}
		date = date.AddDate(0, 0, 1)
	}
}

func (r MonthRule) closestDayOfCertainMonths(start time.Time) time.Time {
	date := start
	for {
		if r.months[date.Month()] {
			day := date.Day()
			day = replaceDayOfMonth(date)
			if r.daysOfMonth[day] {
				return date
			}
		}
		date = date.AddDate(0, 0, 1)
	}
}

func replaceDayOfMonth(date time.Time) int {
	day := date.Day()
	month := date.Month()
	year := date.Year()

	monthDaysCnt := countDaysInMonth(month, year)

	diff := monthDaysCnt - day
	switch diff {
	case 0:
		return 31
	case 1:
		return 30
	}
	return day
}

func countDaysInMonth(month time.Month, year int) int {
	m30 := []time.Month{time.April, time.June, time.September, time.November}

	switch {
	case slices.Contains(m30, month):
		return 30

	case month == time.February:
		if isLeapYear(year) {
			return 29
		}
		return 28

	default:
		return 31
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
