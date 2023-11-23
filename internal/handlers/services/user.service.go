package services

import (
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/models"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) GetAll() ([]models.User, error) {
	users, err := service.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) GetById(id uint) (models.User, error) {
	user, err := service.userRepository.FindById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (service *UserService) Create(user *models.User) error {
	if err := service.userRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (service *UserService) Update(id uint, body models.User) (models.User, error) {
	user, err := service.userRepository.Update(id, body)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (service *UserService) Delete(id uint) error {
	if err := service.userRepository.DeleteById(id); err != nil {
		return err
	}
	return nil
}
