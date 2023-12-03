package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (repo *RoleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	return roles, repo.db.Find(&roles).Error
}

func (repo *RoleRepository) FindById(id uint) (models.Role, error) {
	var role models.Role
	return role, repo.db.First(&role).Error
}

func (repo *RoleRepository) Create(role models.Role) error {
	if err := validation.ValidateStruct(&role); err != nil {
		return err
	}
	return repo.db.Create(&role).Error
}

func (repo *RoleRepository) Update(id uint, roleBody models.Role) (models.Role, error) {
	var existingRole models.Role
	if err := validation.ValidateStruct(roleBody); err != nil {
		return models.Role{}, err
	}

	if err := repo.db.First(&existingRole, id).Error; err != nil {
		return models.Role{}, err
	}

	if err := repo.db.Model(&existingRole).Update(roleBody).Error; err != nil {
		return models.Role{}, err
	}
	return existingRole, nil
}

func (repo *RoleRepository) DeleteById(id uint) error {
	var existingRole models.Role

	if err := repo.db.First(&existingRole, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&existingRole).Error; err != nil {
		return err
	}
	return nil
}
