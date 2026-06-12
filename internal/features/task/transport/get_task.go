package task_transport_http

import (
	"net/http"

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
