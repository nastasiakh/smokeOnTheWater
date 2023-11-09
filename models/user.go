package models

type User struct {
	ID       uint   `json:"id" gorm:"primary key"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
