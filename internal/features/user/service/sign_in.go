package user_service

import (
	"fmt"
	"time"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	tools_jwt "github.com/nikitavaulin/task-manager-golang/internal/core/tools/jwt"
	tools_passwordhasher "github.com/nikitavaulin/task-manager-golang/internal/core/tools/password_hasher"
)

func (s *UserService) SignIn(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", fmt.Errorf("user not found: %w: %w", err, core_errors.ErrNotFound)
	}

	if !tools_passwordhasher.VerifyPassword(password, user.PasswordHash) {
		return "", fmt.Errorf("wrong password: %w", core_errors.ErrInvalidArgument)
	}

	token, err := generateToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func generateToken(username string) (string, error) {
	claims := map[string]any{
		"username":  username,
		"signin_at": time.Now().Unix(),
	}

	token, err := tools_jwt.GenerateToken(claims)
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return token, nil
}
