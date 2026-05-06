package task_service

import (
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

type TaskService struct {
	taskRepo      TaskRepository
	repeatService RepeatTaskSevice
}

type TaskRepository interface {
	CreateTask(task domain.Task) (int64, error)
}

type RepeatTaskSevice interface {
	NextDate(now time.Time, dstart string, repeat string) (string, error)
}

func NewTaskService(taskRepo TaskRepository, repeatService RepeatTaskSevice) *TaskService {
	return &TaskService{
		taskRepo:      taskRepo,
		repeatService: repeatService,
	}
}
