package services

import (
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type RoleService struct {
	roleRepository       *repositories.RoleRepository
	permissionRepository *repositories.PermissionRepository
}

func NewRoleService(roleRepository *repositories.RoleRepository, permissionRepository *repositories.PermissionRepository) *RoleService {
	return &RoleService{roleRepository: roleRepository, permissionRepository: permissionRepository}
}

func (service *RoleService) GetAll() ([]models.Role, error) {
	roles, err := service.roleRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (service *RoleService) GetById(id uint) (models.Role, error) {
	role, err := service.roleRepository.FindById(id)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (service *RoleService) Create(role models.Role) error {
	if err := validation.ValidateStruct(role); err != nil {
		return err
	}
	if err := service.roleRepository.Create(role); err != nil {
		return err
	}
	return nil
}

func (service *RoleService) Update(id uint, body models.Role) (models.Role, error) {
	if err := validation.ValidateStruct(body); err != nil {
		return models.Role{}, err
	}

	role, err := service.roleRepository.Update(id, body)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (service *RoleService) Delete(id uint) error {
	if err := service.roleRepository.DeleteById(id); err != nil {
		return err
	}
	return nil
}
