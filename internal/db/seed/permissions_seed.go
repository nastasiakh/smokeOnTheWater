package seed

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

func CreatePermissionSeed(db *gorm.DB) error {
	permissions := []models.Permission{
		{
			Title: "user.create",
		}, {
			Title: "user.getAll",
		}, {
			Title: "user.getOne",
		}, {
			Title: "user.update",
		}, {
			Title: "user.delete",
		}, {
			Title: "role.create",
		}, {
			Title: "role.getAll",
		}, {
			Title: "role.getOne",
		}, {
			Title: "role.update",
		}, {
			Title: "role.delete",
		},
	}

	for _, permission := range permissions {
		if err := db.Create(&permission).Error; err != nil {
			return err
		}
	}

	return nil
}
