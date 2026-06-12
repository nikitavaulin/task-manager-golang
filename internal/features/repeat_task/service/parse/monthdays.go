package repeat_rule_parser

import (
	"fmt"
	"strconv"
	"strings"

	core_validation "github.com/nikitavaulin/task-manager-golang/internal/core/tools/validation"
)

const (
	lastDayOfMonth       = -1
	beforeLastDayOfMonth = -2
)

type DaysOfMonthSelected struct {
	daysOfMonth     map[int]bool
	isLastDay       bool
	isBeforeLastDay bool
}

func parseDaysOfMonth(daysStr string) (*DaysOfMonthSelected, error) {
	daysOfMonth := make(map[int]bool, 31)
	var (
		isLastDaySelected       bool
		isBeforeLastDaySelected bool
	)

	buf := strings.Split(daysStr, ",")
	if err := core_validation.ValidateIsIntInBounds(len(buf), 1, 33); err != nil {
		return nil, fmt.Errorf("wrong number of days of month in param: %w", err)
	}

	for _, dayStr := range buf {
		day, err := parseDayOfMonth(dayStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse day of month: %w", err)
		}
		if isDayLastOrBeforeLast(day) {
			if day == lastDayOfMonth {
				isLastDaySelected = true
			}
			if day == beforeLastDayOfMonth {
				isBeforeLastDaySelected = true
			}
		} else {
			daysOfMonth[day] = true
		}
	}

	if len(daysOfMonth) == 0 && !isLastDaySelected && !isBeforeLastDaySelected {
		return nil, fmt.Errorf("no days selected")
	}

	return &DaysOfMonthSelected{
		daysOfMonth:     daysOfMonth,
		isLastDay:       isLastDaySelected,
		isBeforeLastDay: isBeforeLastDaySelected,
	}, nil
}

func parseDayOfMonth(dayStr string) (int, error) {
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return -1, fmt.Errorf("failed to convert string to int: %w", err)
	}

	if isDayLastOrBeforeLast(day) {
		return day, nil
	}

	if err := core_validation.ValidateIsIntInBounds(day, 1, 31); err != nil {
		return 0, fmt.Errorf("wrong day of month: %w", err)
	}

	return day, nil
}

func isDayLastOrBeforeLast(day int) bool {
	return day == lastDayOfMonth || day == beforeLastDayOfMonth
}
