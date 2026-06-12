package repository

import "github.com/nikitavaulin/task-manager-golang/internal/core/domain"

type TaskRepository interface {
	CreateTask(task domain.Task) (int64, error)
	UpdateTask(task domain.Task) error
	UpdateTaskDate(taskID int64, newDate string) error
	GetTasks(userID int64, limit int, date, title *string) ([]domain.Task, error)
	GetTask(taskID int64) (domain.Task, error)
	DeleteTask(taskID int64) error
}
