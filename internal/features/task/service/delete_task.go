package task_service

import "fmt"

func (s *TaskService) DeleteTask(taskID int64) error {
	if err := s.taskRepo.DeleteTask(taskID); err != nil {
		return fmt.Errorf("failed to delete task in repo")
	}
	return nil
}
