package task_transport_http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type UpdateTaskRequestDTO struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	Comment    string `json:"comment"`
	Repeat     string `json:"repeat"`
	CategoryID string `json:"category_id"`
	UserID     string `json:"user_id"`
}

func (h *TaskHTTPTransportHandler) UpdateTask(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	var taskRequestDTO UpdateTaskRequestDTO

	if err := core_http_request.Decode(r, &taskRequestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode update task dto")
		return
	}

	task, err := updateTaskDomainFromDTO(taskRequestDTO)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get domain from dto")
		return
	}

	if err := h.taskService.UpdateTask(task); err != nil {
		responseHandler.ErrorResponse(err, "failed to update task")
		return
	}

	responseHandler.EmptyJSONResponse(http.StatusOK)
}

func updateTaskDomainFromDTO(dto UpdateTaskRequestDTO) (domain.Task, error) {
	taskID, err := strconv.ParseInt(dto.ID, 10, 64)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to convert task ID: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	categoryID, err := strconv.ParseInt(dto.CategoryID, 10, 64)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to convert category ID: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	userID, err := strconv.ParseInt(dto.UserID, 10, 64)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to convert user ID: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return *domain.NewTask(
		taskID,
		dto.Title,
		dto.Date,
		dto.Comment,
		dto.Repeat,
		categoryID,
		userID,
	), nil
}
