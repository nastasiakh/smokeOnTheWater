package models

import "smokeOnTheWater/internal/db"

type Role struct {
	ID   uint   `json:"id" gorm:"primary key"`
	Name string `json:"name" gorm:"unique;not null" validate:"required"`
}

func CreateRole(Role *Role) (err error) {
	err = db.DB.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRole(Role *Role, id int) (err error) {
	err = db.DB.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}
