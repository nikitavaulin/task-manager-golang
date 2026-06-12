package auth_service

type AuthService struct {
	appPassword string
}

func NewAuthService(appPassword string) *AuthService {
	return &AuthService{
		appPassword: appPassword,
	}
}
