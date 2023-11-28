package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"smokeOnTheWater/internal/models"
	"time"
)

func MigrateDB(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: generateMigrationID(),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}, &models.Role{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users", "roles").Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}

func generateMigrationID() string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s%s", timestamp, "01")
}
