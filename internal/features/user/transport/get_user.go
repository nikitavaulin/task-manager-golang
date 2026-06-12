package user_transport_http

import (
	"net/http"

	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

func (h *UserTransportHTTP) GetUser(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	userID, err := core_http_request.GetPathValueID(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user ID")
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	responseHandler.JSONResponse(user, http.StatusOK)
}
