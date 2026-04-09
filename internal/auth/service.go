package auth

import (
	"errors"
	"stepik_1/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, name, password string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return " ", errors.New(ErrUserExists)
	}

	user := &user.User{
		Email:    "email",
		Password: "pass",
		Name:     "name",
	}

	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
