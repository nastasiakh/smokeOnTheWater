package models

type Permission struct {
	ID    uint   `json:"id" gorm:"primary key"`
	Title string `json:"title" gorm:"unique;not null" validate:"required"`
}
