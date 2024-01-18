package services

import (
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/models"
)

type PermissionService struct {
	permissionRepository *repositories.PermissionRepository
}

func NewPermissionService(permissionRepository *repositories.PermissionRepository) *PermissionService {
	return &PermissionService{permissionRepository: permissionRepository}
}

func (service *PermissionService) GetAll() ([]models.Permission, error) {
	permissions, err := service.permissionRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
