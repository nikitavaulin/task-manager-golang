package repeat_rule_parser

import (
	"fmt"
	"strconv"
	"strings"

	core_validation "github.com/nikitavaulin/task-manager-golang/internal/core/tools/validation"
)

func parseDaysOfMonth(daysStr string) (map[int]bool, error) {
	daysOfMonth := make(map[int]bool, 31)

	buf := strings.Split(daysStr, ",")
	if err := core_validation.ValidateIsIntInBounds(len(buf), 1, 33); err != nil {
		return nil, fmt.Errorf("wrong number of days of month in param: %w", err)
	}

	for _, dayStr := range buf {
		day, err := parseDayOfMonth(dayStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse day of month: %w", err)
		}
		daysOfMonth[day] = true
	}

	if len(daysOfMonth) == 0 {
		return nil, fmt.Errorf("got empty daysOfMonth map")
	}

	return daysOfMonth, nil
}

func parseDayOfMonth(dayStr string) (int, error) {
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return -1, fmt.Errorf("failed to convert string to int: %w", err)
	}

	if day == -1 {
		return 31, nil
	}

	if day == -2 {
		return 30, nil
	}

	if err := core_validation.ValidateIsIntInBounds(day, 1, 31); err != nil {
		return 0, fmt.Errorf("wrong day of month: %w", err)
	}

	return day, nil
}
