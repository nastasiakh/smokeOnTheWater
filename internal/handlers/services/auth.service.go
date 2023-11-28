package services

import (
	"golang.org/x/crypto/bcrypt"
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/models"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (service *AuthService) Authenticate(email, password string) (*models.User, error) {
	user, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
