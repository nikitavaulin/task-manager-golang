package task_category_service

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (s *TaskCategoryService) UpdateCategoryName(categoryID int64, updatedName string) error {
	var c domain.TaskCategory
	if err := c.ValidateCategoryName(updatedName); err != nil {
		return fmt.Errorf("invalid updated category name: %w", err)
	}
	if err := s.categoryRepository.UpdateCategoryName(categoryID, updatedName); err != nil {
		return fmt.Errorf("failed to update category name: %w", err)
	}
	return nil
}
