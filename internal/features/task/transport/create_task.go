package task_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type CreateTaskRequestDTO struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
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

	task := taskDomainFromDTO(taskRequestDTO)

	taskID, err := h.taskService.CreateTask(task)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task in service")
		return
	}

	response := CreateTaskResponseDTO{taskID}
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func taskDomainFromDTO(dto CreateTaskRequestDTO) domain.Task {
	return *domain.NewTaskUninitialized(
		dto.Title,
		dto.Date,
		dto.Comment,
		dto.Repeat,
	)
}
