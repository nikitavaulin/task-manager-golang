package task_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *TaskRepository) UpdateTask(task domain.Task) error {
	query := `
		UPDATE scheduler
		SET	
			title = $1,
			comment = $2, 
			date = $3, 
			repeat = $4
		WHERE
			id = $5;
	`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Comment,
		task.Date,
		task.Repeat,
		task.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update task in db: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows count: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("task with ID=%d: %w", task.ID, err)
	}

	return nil
}

func (r *TaskRepository) UpdateTaskDate(task domain.Task) error {
	query := `
		UPDATE scheduler
		SET	
			date = $1, 
		WHERE
			id = $2;
	`

	result, err := r.db.Exec(
		query,
		task.Date,
		task.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update task date in db: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows count: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("task with ID=%d: %w", task.ID, err)
	}

	return nil
}
