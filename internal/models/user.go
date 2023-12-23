package models

type User struct {
	ID        uint   `json:"id" gorm:"primary key"`
	FirstName string `json:"firstName" gorm:"column:first_name;not null" validate:"required"`
	LastName  string `json:"lastName" gorm:"column:last_name;not null" validate:"required"`
	Phone     string `json:"phone" gorm:"column:phone;not null" validate:"required"`
	Email     string `json:"email" gorm:"column:email;unique;not null" validate:"required,email"`
	Password  string `json:"password" gorm:"column:password;not null" validate:"required"`
	Roles     []Role `gorm:"many2many:user_role;" json:"roles"`
}
