package service

import (
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

type TaskService interface {
	CreateTask(task domain.Task) (int64, error)
	UpdateTask(task domain.Task) error
	GetTasks(userID int64, limit int, search *string) ([]domain.Task, error)
	GetTask(taskID int64) (domain.Task, error)
	DeleteTask(taskID int64) error
	SetTaskDone(taskID int64) error
}

type TaskCategoryService interface {
	CreateCategory(category domain.TaskCategory) (int64, error)
	GetCategories(userID int64) ([]domain.TaskCategory, error)
	GetCategoryByID(categoryID int64) (domain.TaskCategory, error)
	UpdateCategoryName(categoryID int64, updatedName string) error
	DeleteCategory(categoryID int64) error
}

type RepeatTaskService interface {
	NextDate(now time.Time, dstart string, repeat string) (string, error)
}

type AuthService interface {
	SignIn(password string) (string, error)
}

type UserService interface {
	CreateUser(user domain.User, password string) (int64, error)
	GetUserByID(userID int64) (domain.User, error)
	SignIn(username string, password string) (string, error)
}
