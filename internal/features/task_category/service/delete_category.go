package task_category_service

func (s *TaskCategoryService) DeleteCategory(categoryID int64) error {
	return s.categoryRepository.DeleteCategory(categoryID)
}
