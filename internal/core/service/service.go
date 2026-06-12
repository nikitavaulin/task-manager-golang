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

type RepeatTaskService interface {
	NextDate(now time.Time, dstart string, repeat string) (string, error)
}

type AuthService interface {
	SignIn(password string) (string, error)
}

type UserService interface {
	CreateUser(user domain.User) (int64, error)
	GetUser(userID int64) (domain.User, error)
}
