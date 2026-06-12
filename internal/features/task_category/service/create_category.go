package task_category_service

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (s *TaskCategoryService) CreateCategory(category domain.TaskCategory) (int64, error) {
	if err := category.Validate(); err != nil {
		return 0, fmt.Errorf("invalid category: %w", err)
	}

	id, err := s.categoryRepository.CreateCategory(category)
	if err != nil {
		return 0, fmt.Errorf("failed to create category: %w", err)
	}
	return id, nil
}
