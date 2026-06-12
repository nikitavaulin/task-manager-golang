package user_service

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	tools_passwordhasher "github.com/nikitavaulin/task-manager-golang/internal/core/tools/password_hasher"
)

func (s *UserService) CreateUser(user domain.User, password string) (int64, error) {
	if err := user.Validate(); err != nil {
		return 0, fmt.Errorf("invalid user: %w", err)
	}

	if err := user.ValidatePassword(password); err != nil {
		return 0, fmt.Errorf("invalid password: %w", err)
	}

	passwordHash, err := tools_passwordhasher.HashPassword(password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}
	user.PasswordHash = passwordHash

	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}
