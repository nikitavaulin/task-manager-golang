package task_repository

import "github.com/nikitavaulin/task-manager-golang/internal/core/domain"

type TaskModel struct {
	ID         int64
	Title      string
	Date       string
	Comment    string
	Repeat     string
	CategoryID int64
	UserID     int64
}

func taskDomainFromModel(model TaskModel) domain.Task {
	return *domain.NewTask(
		model.ID,
		model.Title,
		model.Date,
		model.Comment,
		model.Repeat,
		model.CategoryID,
		model.UserID,
	)
}

func tasksDomainFromModels(models []TaskModel) []domain.Task {
	tasks := make([]domain.Task, len(models))
	for i, model := range models {
		task := taskDomainFromModel(model)
		tasks[i] = task
	}
	return tasks
}
