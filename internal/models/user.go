package models

type User struct {
	ID        uint   `json:"id" gorm:"primary key"`
	FirstName string `json:"firstName" gorm:"not null" validate:"required"`
	LastName  string `json:"lastName" gorm:"not null" validate:"required"`
	Email     string `json:"email" gorm:"unique;not null" validate:"required;email"`
	Password  string `json:"password" gorm:"not null" validate:"required"`
}
