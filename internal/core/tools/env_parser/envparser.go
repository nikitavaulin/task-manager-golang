package tools_envparser

import (
	"fmt"
	"os"
)

func GetEnvVarOrDefault(key, defaultValue string) string {
	value, err := GetEnvVar(key)
	if err != nil {
		return defaultValue
	}
	return value
}

func GetEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("failed to get env var with key=%s", key)
	}
	return value, nil
}
