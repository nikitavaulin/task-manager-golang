package repeat_rule_parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	core_validation "github.com/nikitavaulin/task-manager-golang/internal/core/tools/validation"
)

func parseWeekDays(daysStr string) (map[time.Weekday]bool, error) {
	days := make(map[time.Weekday]bool, 7)

	buf := strings.Split(daysStr, ",")
	if err := core_validation.ValidateIsIntInBounds(len(buf), 1, 7); err != nil {
		return nil, fmt.Errorf("wrong number of days: %w", err)
	}

	for _, dayStr := range buf {
		day, err := parseWeekDay(dayStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse week day: %w", err)
		}
		days[day] = true
	}

	if len(days) == 0 {
		return nil, fmt.Errorf("no weekdays got")
	}

	return days, nil
}

func parseWeekDay(dayStr string) (time.Weekday, error) {
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return -1, fmt.Errorf("failed to convert string to int: %w", err)
	}

	if err := core_validation.ValidateIsIntInBounds(day, 1, 7); err != nil {
		return -1, fmt.Errorf("wrong day number: %w", err)
	}

	if day == 7 {
		day = 0
	}

	return time.Weekday(day), nil
}
