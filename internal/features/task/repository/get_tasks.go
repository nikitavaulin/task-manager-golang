package task_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *TaskRepository) GetTasks(userID int64, limit int, date, title *string) ([]domain.Task, error) {
	query, args := getTasksQueryWithArgs(userID, limit, date, title)

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
			&taskModel.CategoryID,
			&taskModel.UserID,
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

func getTasksQueryWithArgs(userID int64, limit int, date, title *string) (string, []any) {
	var (
		query string
		args  []any
	)
	switch {
	case date != nil:
		query = `
			SELECT * FROM tasks
			WHERE user_id = $1 AND date = $2
			LIMIT $3;
		`
		args = []any{userID, *date, limit}

	case title != nil:
		query = `
			SELECT * FROM tasks
			WHERE user_id = $1 AND title LIKE $2 OR comment LIKE $2
			ORDER BY date
			LIMIT $3;
		`
		searchReg := "%" + *title + "%"
		args = []any{userID, searchReg, limit}

	default:
		query = `
			SELECT * FROM tasks
			WHERE user_id = $1
			ORDER BY date
			LIMIT $2;
		`
		args = []any{userID, limit}
	}
	return query, args
}
