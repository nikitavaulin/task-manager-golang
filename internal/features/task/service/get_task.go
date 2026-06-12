package task_service

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (s *TaskService) GetTask(taskID int64) (domain.Task, error) {
	task, err := s.taskRepo.GetTask(taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to get task from repo: %w", err)
	}
	return task, nil
}
