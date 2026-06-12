package task_transport_http

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

type TaskResponseDTO struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func taskDTOFromDomain(task domain.Task) TaskResponseDTO {
	return TaskResponseDTO{
		ID:      fmt.Sprint(task.ID),
		Title:   task.Title,
		Date:    task.Date,
		Comment: task.Comment,
		Repeat:  task.Repeat,
	}
}

func getTasksDTOFromDomain(tasks []domain.Task) []TaskResponseDTO {
	dtos := make([]TaskResponseDTO, len(tasks))
	for i, task := range tasks {
		dto := taskDTOFromDomain(task)
		dtos[i] = dto
	}
	return dtos
}
