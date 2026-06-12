package user_service

import "github.com/nikitavaulin/task-manager-golang/internal/core/repository"

type UserService struct {
	userRepo repository.UserRepositroy
}

func NewUserService(userRepo repository.UserRepositroy) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
