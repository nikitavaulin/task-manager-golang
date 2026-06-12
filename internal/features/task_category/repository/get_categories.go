package task_category_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *TaskCategoryRepository) GetCategories(userID int64) ([]domain.TaskCategory, error) {
	query := `
		SELECT * FROM task_categories
		WHERE user_id = $1;
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories from db: %w", err)
	}
	defer rows.Close()

	var categories []domain.TaskCategory
	for rows.Next() {
		var category domain.TaskCategory

		err := rows.Scan(
			&category.ID,
			&category.CategoryName,
			&category.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan categories from db: %w", err)
		}

		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return categories, nil
}
