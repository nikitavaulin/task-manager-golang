package task_service

import (
	"fmt"
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (s *TaskService) GetTasks(limit int, search *string) ([]domain.Task, error) {
	if limit < 1 {
		return nil, fmt.Errorf("limit value should be > 1")
	}
	var dateSearch, titleSearch *string

	if search != nil {
		dateSearch = tryParseDate(*search)
		if dateSearch == nil {
			titleSearch = search
		}
	}

	tasks, err := s.taskRepo.GetTasks(limit, dateSearch, titleSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks from repo: %w", err)
	}

	return tasks, nil
}

func tryParseDate(searchParam string) *string {
	const dateQueryLayout = "02.01.2006"
	date, err := time.Parse(dateQueryLayout, searchParam)
	if err != nil {
		return nil
	}
	dateStr := date.Format(domain.DateLayout)
	return &dateStr
}
