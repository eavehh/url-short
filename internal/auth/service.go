package auth

import "stepik_1/internal/user"

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (*AuthService) Register(email, name, password string) (string, error) {

}
