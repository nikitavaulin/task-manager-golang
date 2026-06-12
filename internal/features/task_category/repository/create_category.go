package task_category_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *TaskCategoryRepository) CreateCategory(category domain.TaskCategory) (int64, error) {
	query := `
		INSERT INTO task_categories
			(category_name, user_id)
		VALUES
			($1, $2);
	`

	result, err := r.db.Exec(
		query,
		category.CategoryName,
		category.UserID,
	)
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to create category in db: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to get category ID: %w", err)
	}

	return id, nil
}
