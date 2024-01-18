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
	return roles, repo.db.Preload("Permissions").Find(&roles).Error
}

func (repo *RoleRepository) FindById(id uint) (models.Role, error) {
	var role models.Role
	return role, repo.db.Preload("Permissions").First(&role, id).Error
}

func (repo *RoleRepository) FindByName(name string) (models.Role, error) {
	var role models.Role
	if err := repo.db.Where("name = ?", name).First(&role).Error; err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (repo *RoleRepository) Create(role models.Role) error {
	if err := validation.ValidateStruct(&role); err != nil {
		return err
	}
	tx := repo.db.Begin()

	if err := repo.db.Create(&role).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, permission := range role.Permissions {
		if err := tx.Model(&role).Association("Permissions").Append(permission).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (repo *RoleRepository) Update(id uint, roleBody models.Role) (models.Role, error) {
	var existingRole models.Role

	if err := validation.ValidateStruct(roleBody); err != nil {
		return models.Role{}, err
	}

	existingRole, err := repo.FindById(id)
	if err != nil {
		return models.Role{}, err
	}

	tx := repo.db.Begin()

	if err = tx.Model(&existingRole).Updates(models.Role{
		Name: roleBody.Name,
	}).Error; err != nil {
		tx.Rollback()
		return models.Role{}, err
	}

	if err := tx.Model(&existingRole).Association("Permissions").Replace(roleBody.Permissions).Error; err != nil {
		tx.Rollback()
		return models.Role{}, err
	}

	tx.Commit()
	return existingRole, nil
}

func (repo *RoleRepository) DeleteById(id uint) error {
	tx := repo.db.Begin()

	var existingRole models.Role
	if err := tx.First(&existingRole, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&existingRole).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
