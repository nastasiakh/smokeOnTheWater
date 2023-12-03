package tables

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func MigrateUsersTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}).Error
}

func RollbackUsersTable(db *gorm.DB) error {
	return db.DropTableIfExists("users").Error
}
