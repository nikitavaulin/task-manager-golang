package core_http_server

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config *HTTPServerConfig
}

func NewHTTPServer(config *HTTPServerConfig) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
	}
}

func (h *HTTPServer) RegisterRouters(routers ...*Router) {
	for _, router := range routers {
		h.mux.Handle("/", router)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    h.config.Address,
		Handler: h.mux,
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- server.ListenAndServe()
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("listen and serve: %w", err)
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), h.config.ShutdownDuration)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("failed to shutdown HTTP server correctly: %w", err)
		}
	}

	return nil
}
