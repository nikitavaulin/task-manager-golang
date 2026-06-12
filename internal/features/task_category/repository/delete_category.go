package task_category_repository

import (
	"fmt"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func (r *TaskCategoryRepository) DeleteCategory(categoryID int64) error {
	query := `
		DELETE FROM task_categories
		WHERE id = $1;
	`

	result, err := r.db.Exec(query, categoryID)
	if err != nil {
		return fmt.Errorf("failed to delete category in db: %w", err)
	}

	deleted, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get deleted category count: %w", err)
	}

	if deleted == 0 {
		return fmt.Errorf("category with ID=%d: %w", categoryID, core_errors.ErrNotFound)
	}

	return nil
}
