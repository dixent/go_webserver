package services

import (
	"go_webserver/internal/auth/entities"
	"go_webserver/internal/auth/repositories"
	"go_webserver/pkg/jwt"
)

type AuthenticateService struct {
	userRepo repositories.UserRepository
}

func NewAuthenticateService(userRepo repositories.UserRepository) *AuthenticateService {
	return &AuthenticateService{userRepo}
}

func (s *AuthenticateService) Authenticate(auth entities.Auth) string {
	user := s.userRepo.GetByAuth(&auth)

	if user.Id == 0 {
		return ""
	}

	tokenValidator := jwt.NewTokenValidator()
	token, err := tokenValidator.EncryptToken(&user)

	if err != nil {
		return ""
	}

	return token
}
