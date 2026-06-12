package task_category_repository

import (
	"fmt"
)

func (r *TaskCategoryRepository) UpdateCategoryName(categoryID int64, updatedName string) error {
	query := `
		UPDATE task_categories
		SET	
			category_name = $1 
		WHERE
			id = $2;
	`

	result, err := r.db.Exec(
		query,
		updatedName,
		categoryID,
	)
	if err != nil {
		return fmt.Errorf("failed to update category name in db: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows count: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("category with ID=%d: %w", categoryID, err)
	}

	return nil
}
