package task_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *TaskRepository) CreateTask(task domain.Task) (int64, error) {
	query := `
		INSERT INTO scheduler
			(title, comment, date, repeat)
		VALUES
			($1, $2, $3, $4);
	`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Comment,
		task.Date,
		task.Repeat,
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
