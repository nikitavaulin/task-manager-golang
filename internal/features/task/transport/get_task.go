package task_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type GetTaskResponseDTO TaskResponseDTO

func (h *TaskHTTPTransportHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	taskID, err := getTaskID(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task ID")
		return
	}

	task, err := h.taskService.GetTask(taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}

	taskDTO := GetTaskResponseDTO(taskDTOFromDomain(task))
	responseHandler.JSONResponse(taskDTO, http.StatusOK)
}

func getTaskID(r *http.Request) (int64, error) {
	id, err := core_http_request.GetIntQueryParam("id", r)
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to get ID query param")
	}
	if id == nil {
		return domain.UninitializedID, fmt.Errorf("ID cannot be empty: %w", core_errors.ErrInvalidArgument)
	}
	return int64(*id), err
}
