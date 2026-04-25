package core_http_server

import (
	"os"
	"time"
)

type HTTPServerConfig struct {
	Address          string
	ShutdownDuration time.Duration
}

func NewHTTPServerConfig() *HTTPServerConfig {

	return &HTTPServerConfig{
		Address:          ParseServerPort(),
		ShutdownDuration: ParseShutdownDuration(),
	}
}

func ParseServerPort() string {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		return ":7540"
	}
	return port
}

func ParseShutdownDuration() time.Duration {
	defaultDuration := 5 * time.Second
	durationString := os.Getenv("HTTP_SHUTDOWN_DURATION")

	shutdownDuration, err := time.ParseDuration(durationString)
	if err != nil {
		return defaultDuration
	}

	return shutdownDuration
}
