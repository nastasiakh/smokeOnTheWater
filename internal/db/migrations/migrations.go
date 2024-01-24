package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"smokeOnTheWater/internal/db/migrations/tables"
)

func MigrateDB(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: GenerateMigrationID(),
			Migrate: func(tx *gorm.DB) error {
				if err := tables.MigrateUsersTable(tx); err != nil {
					return err
				}
				if err := tables.MigrateRolesTable(tx); err != nil {
					return err
				}
				if err := tables.MigratePermissionTable(tx); err != nil {
					return err
				}
				if err := tables.MigrateCategoriesTable(tx); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tables.RollbackUsersTable(tx); err != nil {
					return err
				}
				if err := tables.RollbackRolesTable(tx); err != nil {
					return err
				}
				if err := tables.RollbackPermissionTable(tx); err != nil {
					return err
				}
				if err := tables.RollbackCategoriesTable(tx); err != nil {
					return err
				}
				return nil
			},
		},
	})

	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}
