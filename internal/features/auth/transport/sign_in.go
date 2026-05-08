package auth_transport_http

import (
	"net/http"

	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type SignInRequestDTO struct {
	Password string `json:"password"`
}

type SignInResponseDTO struct {
	Token string `json:"token"`
}

func (h *AuthHTTPTrasnportHandler) SignIn(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	var requestDTO SignInRequestDTO
	if err := core_http_request.Decode(r, &requestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode sign in request")
		return
	}

	token, err := h.authService.SignIn(requestDTO.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to sign in")
		return
	}

	responseDTO := SignInResponseDTO{token}
	responseHandler.JSONResponse(responseDTO, http.StatusOK)
}
