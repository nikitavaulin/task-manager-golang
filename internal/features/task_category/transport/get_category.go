package task_category_transport_http

import (
	"net/http"

	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

func (h *TaskCategoryTransportHTTP) GetCategory(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	categoryID, err := core_http_request.GetPathValueID(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category ID")
		return
	}

	user, err := h.categoryService.GetCategoryByID(categoryID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category")
		return
	}

	responseHandler.JSONResponse(user, http.StatusOK)
}
