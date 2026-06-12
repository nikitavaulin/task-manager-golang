package core_http_middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(handler http.Handler, middleware ...Middleware) http.Handler {
	if len(middleware) == 0 {
		return handler
	}
	for idx := len(middleware) - 1; idx >= 0; idx-- {
		handler = middleware[idx](handler)
	}
	return handler
}
