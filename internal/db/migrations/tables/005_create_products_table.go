package tables

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func MigrateProductsTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Product{}).Error
}

func RollbackProductsTable(db *gorm.DB) error {

	return db.DropTableIfExists("products").Error
}
