package task_service

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (s *TaskService) UpdateTask(task domain.Task) error {
	if err := task.Validate(); err != nil {
		return fmt.Errorf("invalid task: %w", err)
	}

	if err := s.checkDateAndRepeat(&task); err != nil {
		return fmt.Errorf("invalid date or repeat: %w", err)
	}

	if err := s.taskRepo.UpdateTask(task); err != nil {
		return fmt.Errorf("failed to update task in repo: %w", err)
	}

	return nil
}
