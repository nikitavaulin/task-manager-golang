package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Auth       bool
	Middleware []core_http_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	h := core_http_middleware.ChainMiddleware(r.Handler, r.Middleware...)
	if r.Auth {
		h = core_http_middleware.Auth()(h)
	}
	return h
}
