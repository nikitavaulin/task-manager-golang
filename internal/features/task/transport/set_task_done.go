package task_transport_http

import (
	"net/http"

	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

func (h *TaskHTTPTransportHandler) SetTaskDone(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	taskID, err := getTaskID(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task ID")
		return
	}

	if err := h.taskService.SetTaskDone(taskID); err != nil {
		responseHandler.ErrorResponse(err, "failed to set task done")
		return
	}

	responseHandler.EmptyJSONResponse(http.StatusOK)
}
