package tables

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func MigratePermissionTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Permission{}).Error
}

func RollbackPermissionTable(db *gorm.DB) error {
	return db.DropTableIfExists("permissions").Error
}
