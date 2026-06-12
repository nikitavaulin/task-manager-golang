package task_category_service

import "github.com/nikitavaulin/task-manager-golang/internal/core/domain"

func (s *TaskCategoryService) GetCategories(userID int64) ([]domain.TaskCategory, error) {
	return s.categoryRepository.GetCategories(userID)
}
