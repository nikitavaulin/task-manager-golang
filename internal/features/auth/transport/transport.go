package auth_transport_http

import (
	"net/http"

	core_http_server "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/server"
)

type AuthHTTPTrasnportHandler struct {
	authService AuthService
}

type AuthService interface {
	SignIn(password string) (string, error)
}

func NewAuthHTTPTrasnportHandler(authService AuthService) *AuthHTTPTrasnportHandler {
	return &AuthHTTPTrasnportHandler{
		authService: authService,
	}
}

func (h *AuthHTTPTrasnportHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/signin",
			Handler: h.SignIn,
		},
	}
}
