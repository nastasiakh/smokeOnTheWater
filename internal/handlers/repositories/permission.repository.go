package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (repo *PermissionRepository) FindAll() ([]models.Permission, error) {
	var permissions []models.Permission
	return permissions, repo.db.Find(&permissions).Error
}
