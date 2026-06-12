package task_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *TaskRepository) CreateTask(task domain.Task) (int64, error) {
	query := `
		INSERT INTO tasks
			(title, comment, date, repeat, category_id, user_id)
		VALUES
			($1, $2, $3, $4, $5, $6);
	`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Comment,
		task.Date,
		task.Repeat,
		task.CategoryID,
		task.UserID,
	)
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to create task in db: %w", err)
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to get taskID: %w", err)
	}

	return taskID, nil
}
