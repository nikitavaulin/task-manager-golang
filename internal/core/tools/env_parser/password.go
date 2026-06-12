package tools_envparser

import "fmt"

func GetAppPassword() (string, error) {
	appPassword, err := GetEnvVar("TODO_PASSWORD")
	if err != nil {
		return "", fmt.Errorf("failed to get app password: %w", err)
	}
	return appPassword, nil
}
