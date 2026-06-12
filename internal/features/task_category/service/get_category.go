package task_category_service

import "github.com/nikitavaulin/task-manager-golang/internal/core/domain"

func (s *TaskCategoryService) GetCategoryByID(categoryID int64) (domain.TaskCategory, error) {
	return s.categoryRepository.GetCategoryByID(categoryID)
}
