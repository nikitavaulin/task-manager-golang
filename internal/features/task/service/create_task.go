package task_service

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (s *TaskService) CreateTask(task domain.Task) (int64, error) {
	if err := task.Validate(); err != nil {
		return domain.UninitializedID, fmt.Errorf("invalid task: %w", err)
	}

	if err := s.checkDateAndRepeat(&task); err != nil {
		return domain.UninitializedID, fmt.Errorf("invalid date or repeat: %w", err)
	}

	taskID, err := s.taskRepo.CreateTask(task)
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to create task in repo: %w", err)
	}

	return taskID, nil
}
