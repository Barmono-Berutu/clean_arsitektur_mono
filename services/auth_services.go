package services

import (
	"clean/domain"
	"clean/security"
	"errors"
)

type AuthService struct {
	userRepo domain.UserRepository
}

func NewAuthService(repo domain.UserRepository) domain.UserUsecase {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Register(user *domain.User) error {
	hash, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(username string, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", err
	}
	if !security.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	token, err := security.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
