package seed

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/db/seed"
)

func SeedData(db *gorm.DB) error {
	if err := seed.CreatePermissionSeed(db); err != nil {
		return err
	}

	return nil
}
