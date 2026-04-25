package core_http_server

import (
	"fmt"
	"net/http"
)

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

func (r *Router) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, route.Handler)
	}
}

func (r *Router) RegisterFileServer(pattern, dirPath string) {
	r.Handle(pattern, http.FileServer(http.Dir(dirPath)))
}
