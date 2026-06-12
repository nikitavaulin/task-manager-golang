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

type TaskCategoryRepository interface {
	CreateCategory(category domain.TaskCategory) (int64, error)
	GetCategories(userID int64) ([]domain.TaskCategory, error)
	GetCategoryByID(categoryID int64) (domain.TaskCategory, error)
	UpdateCategoryName(categoryID int64, updatedName string) error
	DeleteCategory(categoryID int64) error
}

type UserRepositroy interface {
	CreateUser(user domain.User) (int64, error)
	GetUserByID(userID int64) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)
}
