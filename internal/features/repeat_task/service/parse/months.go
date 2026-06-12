package repeat_rule_parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	core_validation "github.com/nikitavaulin/task-manager-golang/internal/core/tools/validation"
)

func parseMonths(monthsStr string) (map[time.Month]bool, error) {
	months := make(map[time.Month]bool, 12)

	buf := strings.Split(monthsStr, ",")
	if err := core_validation.ValidateIsIntInBounds(len(buf), 1, 12); err != nil {
		return nil, fmt.Errorf("wrong number of months: %w", err)
	}

	for _, monthStr := range buf {
		month, err := parseMonth(monthStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse month: %w", err)
		}
		months[month] = true
	}

	if len(months) == 0 {
		return nil, fmt.Errorf("got empty months map")
	}

	return months, nil
}

func parseMonth(monthStr string) (time.Month, error) {
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return -1, fmt.Errorf("failed to convert string to int: %w", err)
	}

	if err := core_validation.ValidateIsIntInBounds(month, 1, 12); err != nil {
		return -1, fmt.Errorf("wrong month number: %w", err)
	}

	return time.Month(month), nil
}
