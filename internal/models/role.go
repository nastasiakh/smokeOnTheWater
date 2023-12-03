package models

type Role struct {
	ID   uint   `json:"id" gorm:"primary key"`
	Name string `json:"name" gorm:"unique;not null" validate:"required"`
}
