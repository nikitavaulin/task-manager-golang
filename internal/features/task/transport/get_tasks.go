package task_transport_http

import (
	"net/http"

	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type GetTasksResponseDTO struct {
	Tasks []TaskResponseDTO `json:"tasks"`
}

const MaxTasksReturnCount = 50

func (h *TaskHTTPTransportHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	searchParam := core_http_request.GetStringQueryParam("search", r)

	tasks, err := h.taskService.GetTasks(MaxTasksReturnCount, searchParam)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks from service")
		return
	}

	tasksDTO := GetTasksResponseDTO{Tasks: getTasksDTOFromDomain(tasks)}
	responseHandler.JSONResponse(tasksDTO, http.StatusOK)
}
