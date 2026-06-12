package user_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type IDResponse struct {
	ID string `json:"id"`
}

func (h *UserTransportHTTP) CreateUser(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	var user domain.User
	if err := core_http_request.Decode(r, &user); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode user")
		return
	}

	userID, err := h.userService.CreateUser(user)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	responseHandler.JSONResponse(IDResponse{fmt.Sprint(userID)}, http.StatusCreated)
}
