package user_service

import "github.com/nikitavaulin/task-manager-golang/internal/core/domain"

func (s *UserService) GetUserByID(userID int64) (domain.User, error) {
	return s.userRepo.GetUserByID(userID)
}
