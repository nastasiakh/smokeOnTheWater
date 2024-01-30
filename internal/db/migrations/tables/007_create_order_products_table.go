package tables

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func MigrateOrderProductsTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.OrderProduct{}).Error
}

func RollbackOrderProductsTable(db *gorm.DB) error {
	return db.DropTableIfExists("orderProducts").Error
}
