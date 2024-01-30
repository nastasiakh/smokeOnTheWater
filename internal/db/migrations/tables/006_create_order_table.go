package tables

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func MigrateOrdersTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Order{}).Error
}

func RollbackOrdersTable(db *gorm.DB) error {
	return db.DropTableIfExists("orders").Error
}
