package task_category_service

import "github.com/nikitavaulin/task-manager-golang/internal/core/repository"

type TaskCategoryService struct {
	categoryRepository repository.TaskCategoryRepository
}

func NewTaskCategoryService(categoryRepository repository.TaskCategoryRepository) *TaskCategoryService {
	return &TaskCategoryService{
		categoryRepository: categoryRepository,
	}
}
