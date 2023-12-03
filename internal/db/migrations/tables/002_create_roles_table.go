package tables

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func MigrateRolesTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Role{}).Error
}

func RollbackRolesTable(db *gorm.DB) error {
	return db.DropTableIfExists("roles").Error
}
