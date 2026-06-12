package task_category_transport_http

import (
	"net/http"

	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

func (h *TaskCategoryTransportHTTP) GetCategories(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	userID, err := core_http_request.GetIDQueryParamMust("user", r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user ID")
		return
	}

	categories, err := h.categoryService.GetCategories(userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get categories")
		return
	}

	responseHandler.JSONResponse(categories, http.StatusOK)
}
