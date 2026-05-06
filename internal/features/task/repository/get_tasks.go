package task_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *TaskRepository) GetTasks(limit int, date, title *string) ([]domain.Task, error) {
	query, args := getTasksQueryWithArgs(limit, date, title)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks from db: %w", err)
	}
	defer rows.Close()

	var tasksModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel

		err := rows.Scan(
			&taskModel.ID,
			&taskModel.Title,
			&taskModel.Comment,
			&taskModel.Date,
			&taskModel.Repeat,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tasks from db: %w", err)
		}

		tasksModels = append(tasksModels, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	tasks := tasksDomainFromModels(tasksModels)
	return tasks, nil
}

func getTasksQueryWithArgs(limit int, date, title *string) (string, []any) {
	var (
		query string
		args  []any
	)
	switch {
	case date != nil:
		query = `
			SELECT * FROM scheduler
			WHERE date = $1
			LIMIT $2;
		`
		args = []any{*date, limit}

	case title != nil:
		query = `
			SELECT * FROM scheduler
			WHERE title LIKE $1 OR comment LIKE $1
			ORDER BY date
			LIMIT $2;
		`
		searchReg := "%" + *title + "%"
		args = []any{searchReg, limit}

	default:
		query = `
			SELECT * FROM scheduler
			ORDER BY date
			LIMIT $1;
		`
		args = []any{limit}
	}
	return query, args
}
