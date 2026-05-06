package task_repository

import "fmt"

func (r *TaskRepository) DeleteTask(taskID int64) error {
	query := `
		DELETE FROM scheduler
		WHERE id = $1;
	`

	result, err := r.db.Exec(query, taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task in db: %w", err)
	}

	deleted, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get deleted tasks count: %w", err)
	}

	if deleted == 0 {
		return fmt.Errorf("task with ID=%d: %w", taskID, err)
	}

	return nil
}
