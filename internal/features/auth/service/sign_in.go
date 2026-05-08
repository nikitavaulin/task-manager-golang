package auth_service

import (
	"fmt"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	tools_envparser "github.com/nikitavaulin/task-manager-golang/internal/core/tools/env_parser"
	tools_jwt "github.com/nikitavaulin/task-manager-golang/internal/core/tools/jwt"
	tools_passwordhasher "github.com/nikitavaulin/task-manager-golang/internal/core/tools/password_hasher"
)

func (s *AuthService) SignIn(password string) (string, error) {
	if err := validatePassword(password); err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	appPassword, err := tools_envparser.GetAppPassword()
	if err != nil {
		return "", fmt.Errorf("SignIn: %w", err)
	}

	if !verifyPassword(appPassword, password) {
		return "", fmt.Errorf("wrong password: %w", core_errors.ErrInvalidArgument)
	}

	token, err := generateToken(password)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func validatePassword(password string) error {
	if len(password) == 0 {
		return fmt.Errorf("password is empty: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func verifyPassword(appPassword, userPassword string) bool {
	return appPassword == userPassword
}

func generateToken(password string) (string, error) {
	passwordHash, err := tools_passwordhasher.HashPassword(password)
	if err != nil {
		return "", fmt.Errorf("generateToken: %w", err)
	}

	claims := map[string]any{
		"password": passwordHash,
	}

	token, err := tools_jwt.GenerateToken(claims)
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return token, nil
}
