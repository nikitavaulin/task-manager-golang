package task_transport_http

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

type TaskResponseDTO struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	Comment    string `json:"comment"`
	Repeat     string `json:"repeat"`
	CategoryID string `json:"category_id"`
	UserID     string `json:"user_id"`
}

func taskDTOFromDomain(task domain.Task) TaskResponseDTO {
	return TaskResponseDTO{
		ID:         fmt.Sprint(task.ID),
		Title:      task.Title,
		Date:       task.Date,
		Comment:    task.Comment,
		Repeat:     task.Repeat,
		CategoryID: fmt.Sprint(task.CategoryID),
		UserID:     fmt.Sprint(task.UserID),
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
