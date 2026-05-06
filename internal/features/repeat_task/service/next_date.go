package repeat_service

import (
	"fmt"
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	repeat_rule_parser "github.com/nikitavaulin/task-manager-golang/internal/features/repeat_task/service/parse"
)

func (s *RepeatTaskSevice) NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", fmt.Errorf("failed to calc next date: %w", core_errors.ErrNoRepeatRule)
	}

	start, err := time.Parse(domain.DateLayout, dstart)
	if err != nil {
		return "", fmt.Errorf("NextDate: failed to parse start date: %w", err)
	}

	repeatRule, err := repeat_rule_parser.ParseRepeatRule(repeat)
	if err != nil {
		return "", fmt.Errorf("failed to parse repeat rule: %w", err)
	}

	nextDate := repeatRule.CalcNextDate(now, start)

	return nextDate.Format(domain.DateLayout), nil
}
