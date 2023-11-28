package models

type User struct {
	ID        uint   `json:"id" gorm:"primary key"`
	FirstName string `json:"firstName" gorm:"column:first_name;not null" validate:"required"`
	LastName  string `json:"lastName" gorm:"column:last_name;not null" validate:"required"`
	Email     string `json:"email" gorm:"column:email;unique;not null" validate:"required,email"`
	Password  string `json:"password" gorm:"column:password;not null" validate:"required"`
}
