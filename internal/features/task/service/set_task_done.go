package task_service

import "fmt"

func (s *TaskService) SetTaskDone(taskID int64) error {
	task, err := s.taskRepo.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	if len(task.Repeat) == 0 {
		return s.taskRepo.DeleteTask(taskID)
	}

	now := todayDate()
	nextDate, err := s.repeatService.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return fmt.Errorf("failed to calc next date: %w", err)
	}

	return s.taskRepo.UpdateTaskDate(taskID, nextDate)
}
