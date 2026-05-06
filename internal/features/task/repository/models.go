package task_repository

import "github.com/nikitavaulin/task-manager-golang/internal/core/domain"

type TaskModel struct {
	ID      int64
	Title   string
	Date    string
	Comment string
	Repeat  string
}

func taskDomainFromModel(model TaskModel) domain.Task {
	return *domain.NewTask(
		model.ID,
		model.Title,
		model.Date,
		model.Comment,
		model.Repeat,
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
