package task_service

import (
	"errors"
	"fmt"
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func (s *TaskService) CreateTask(task domain.Task) (int64, error) {
	if err := task.Validate(); err != nil {
		return domain.UninitializedID, fmt.Errorf("invalid task: %w", err)
	}

	if err := s.checkDate(&task); err != nil {
		return domain.UninitializedID, fmt.Errorf("invalid date or repeat: %w", err)
	}

	taskID, err := s.taskRepo.CreateTask(task)
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to create task in repo: %w", err)
	}

	return taskID, nil
}

func (s *TaskService) checkDate(task *domain.Task) error {
	now := todayDate()

	date, err := time.Parse(domain.DateLayout, task.Date)
	if err != nil {
		return fmt.Errorf("incorrect date format: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	nextDate, err := s.repeatService.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		if !errors.Is(err, core_errors.ErrNoRepeatRule) {
			return fmt.Errorf("failed to get nextDate: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}

	if now.After(date) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(domain.DateLayout)
		} else {
			task.Date = nextDate
		}
	}

	return nil
}

func todayDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
