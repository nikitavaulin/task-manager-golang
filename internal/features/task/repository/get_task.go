package task_repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func (r *TaskRepository) GetTask(taskID int64) (domain.Task, error) {
	query := `
		SELECT * FROM scheduler
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, taskID)

	var taskModel TaskModel

	err := row.Scan(
		&taskModel.ID,
		&taskModel.Title,
		&taskModel.Comment,
		&taskModel.Date,
		&taskModel.Repeat,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task with ID=%d: %v: %w", taskID, err, core_errors.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("failed to get task from db: %w", err)
	}

	task := taskDomainFromModel(taskModel)
	return task, nil
}
