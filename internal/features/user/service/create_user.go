package user_service

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (s *UserService) CreateUser(user domain.User) (int64, error) {
	if err := user.Validate(); err != nil {
		return 0, fmt.Errorf("invalid user: %w", err)
	}
	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}
