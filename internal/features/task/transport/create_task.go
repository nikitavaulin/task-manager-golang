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

type CreateTaskRequestDTO struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	Comment    string `json:"comment"`
	Repeat     string `json:"repeat"`
	CategoryID string `json:"category_id"`
	UserID     string `json:"user_id"`
}

type CreateTaskResponseDTO struct {
	ID int64 `json:"id"`
}

func (h *TaskHTTPTransportHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	var taskRequestDTO CreateTaskRequestDTO

	if err := core_http_request.Decode(r, &taskRequestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode create task dto")
		return
	}

	task, err := taskDomainFromDTO(taskRequestDTO)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to convert dto to domain task")
		return
	}

	taskID, err := h.taskService.CreateTask(task)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task in service")
		return
	}

	response := CreateTaskResponseDTO{taskID}
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func taskDomainFromDTO(dto CreateTaskRequestDTO) (domain.Task, error) {
	categoryID, err := strconv.ParseInt(dto.CategoryID, 10, 64)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to convert category ID: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	userID, err := strconv.ParseInt(dto.UserID, 10, 64)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to convert user ID: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return *domain.NewTaskUninitialized(
		dto.Title,
		dto.Date,
		dto.Comment,
		dto.Repeat,
		categoryID,
		userID,
	), nil
}
