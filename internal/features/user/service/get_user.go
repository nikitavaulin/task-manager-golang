package user_service

import "github.com/nikitavaulin/task-manager-golang/internal/core/domain"

func (s *UserService) GetUser(userID int64) (domain.User, error) {
	return s.userRepo.GetUser(userID)
}
