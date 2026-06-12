package task_category_repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func (r *TaskCategoryRepository) GetCategoryByID(categoryID int64) (domain.TaskCategory, error) {
	query := `
		SELECT * FROM task_categories
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, categoryID)

	var category domain.TaskCategory

	err := row.Scan(
		&category.ID,
		&category.CategoryName,
		&category.UserID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.TaskCategory{}, fmt.Errorf("category with ID=%d: %v: %w", categoryID, err, core_errors.ErrNotFound)
		}
		return domain.TaskCategory{}, fmt.Errorf("failed to get category from db: %w", err)
	}

	return category, nil
}
