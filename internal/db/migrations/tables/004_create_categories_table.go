package tables

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func MigrateCategoriesTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Category{}).Error
}

func RollbackCategoriesTable(db *gorm.DB) error {
	return db.DropTableIfExists("categories").Error
}
