package task_service

import (
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/core/repository"
)

type TaskService struct {
	taskRepo      repository.TaskRepository
	repeatService RepeatTaskSevice
}

type RepeatTaskSevice interface {
	NextDate(now time.Time, dstart string, repeat string) (string, error)
}

func NewTaskService(taskRepo repository.TaskRepository, repeatService RepeatTaskSevice) *TaskService {
	return &TaskService{
		taskRepo:      taskRepo,
		repeatService: repeatService,
	}
}
