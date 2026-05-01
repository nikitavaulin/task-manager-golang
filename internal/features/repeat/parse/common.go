package repeat_rule_parser

import (
	"fmt"
	"strconv"
	"strings"

	repeat_rules "github.com/nikitavaulin/task-manager-golang/internal/features/repeat/rules"
)

func ParseRepeatRule(repeat string) (repeat_rules.RepeatRule, error) {
	buf := strings.Split(repeat, " ")

	if len(buf) < 1 {
		return nil, fmt.Errorf("failed to parse repeat rule")
	}

	argsCount := len(buf) - 1

	switch buf[0] {
	case "y":
		if argsCount > 0 {
			return nil, fmt.Errorf("param 'y' cannot have any arguments, got: %d", argsCount)
		}

		return repeat_rules.NewYearRule(), nil

	case "d":
		if argsCount != 1 {
			return nil, fmt.Errorf("param 'd' should have 1 argument, got: %d", argsCount)
		}

		daysCount, err := strconv.Atoi(buf[1])
		if err != nil {
			return nil, fmt.Errorf("could not convert daysCount string to int: %w", err)
		}

		return repeat_rules.NewDayRule(daysCount)

	case "w":
		if argsCount != 1 {
			return nil, fmt.Errorf("param 'w' should have 1 argument, got: %d", argsCount)
		}

		daysOfWeek, err := parseWeekDays(buf[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse days of week: %w", err)
		}

		return repeat_rules.NewWeekRule(daysOfWeek)

	case "m":
		if argsCount != 2 && argsCount != 1 {
			return nil, fmt.Errorf("param 'm' should have 1 or 2 argument(s), got: %d", argsCount)
		}

		days, err := parseDaysOfMonth(buf[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse days of month: %w", err)
		}

		if argsCount == 2 {
			months, err := parseMonths(buf[2])
			if err != nil {
				return nil, fmt.Errorf("failed to parse months: %w", err)
			}
			repeat_rules.NewMonthRule(days, months)
		}

		return repeat_rules.NewMonthRule(days, nil)
	}

	return nil, fmt.Errorf("unknown time measurement unit")
}
