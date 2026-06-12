package user_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/service"
	core_http_server "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/server"
)

type UserTransportHTTP struct {
	userService service.UserService
}

func NewUserTransportHTTP(userService service.UserService) *UserTransportHTTP {
	return &UserTransportHTTP{
		userService: userService,
	}
}

func (h *UserTransportHTTP) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/api/users/{id}",
			Handler: h.GetUser,
			Auth:    true,
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/auth/register",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/auth/login",
			Handler: h.SignIn,
		},
	}
}
