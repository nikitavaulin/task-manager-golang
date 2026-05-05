package repeat_service

import (
	"errors"
	"fmt"
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/features/repeat_task"
	repeat_rule_parser "github.com/nikitavaulin/task-manager-golang/internal/features/repeat_task/service/parse"
)

var ErrNoRepeatRule = errors.New("repeat rule is empty")

func (s *RepeatTaskSevice) NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", fmt.Errorf("failed to calc next date: %w", ErrNoRepeatRule)
	}

	start, err := time.Parse(repeat_task.DateLayout, dstart)
	if err != nil {
		return "", fmt.Errorf("NextDate: failed to parse start date: %w", err)
	}

	repeatRule, err := repeat_rule_parser.ParseRepeatRule(repeat)
	if err != nil {
		return "", fmt.Errorf("failed to parse repeat rule: %w", err)
	}

	nextDate := repeatRule.CalcNextDate(now, start)

	return nextDate.Format(repeat_task.DateLayout), nil
}
