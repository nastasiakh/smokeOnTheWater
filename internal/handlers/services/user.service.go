package services

import (
	"golang.org/x/crypto/bcrypt"
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type UserService struct {
	userRepository *repositories.UserRepository
	roleRepository *repositories.RoleRepository
}

func NewUserService(userRepository *repositories.UserRepository, roleRepository *repositories.RoleRepository) *UserService {
	return &UserService{userRepository: userRepository, roleRepository: roleRepository}
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

func (service *UserService) GetByEmail(email string) (*models.User, error) {
	user, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (service *UserService) Create(user *models.User) (*models.User, error) {

	if err := validation.ValidateStruct(*user); err != nil {
		return nil, err
	}

	// todo create helper for hasher
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// todo check isEmpty method or similar
	if len(user.Roles) == 0 {
		clientRole, err := service.roleRepository.FindByName("client")
		if err != nil {
			return nil, nil
		}
		user.Roles = append(user.Roles, clientRole)
	}

	createdUser, err := service.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (service *UserService) Update(id uint, body *models.User) (models.User, error) {
	if err := validation.ValidateStruct(body); err != nil {
		return models.User{}, err
	}
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
