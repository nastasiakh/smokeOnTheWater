package models

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primary key"`
	TotalAmount  float64   `json:"totalAmount" gorm:"column:totalAmount;not null" validate:"required"`
	DateCreated  time.Time `json:"dateCreated" gorm:"column:dateCreated;not null" validate:"required"`
	DateModified time.Time `json:"dateModified" gorm:"column:dateModified;not null" validate:"required"`
	Status       string    `json:"status" gorm:"column:status;not null" validate:"required"`
	FirstName    string    `json:"firstName" gorm:"column:firstName;not null" validate:"required"`
	LastName     string    `json:"lastName" gorm:"column:lastName;not null" validate:"required"`
	Phone        string    `json:"phone" gorm:"column:phone;not null" validate:"required"`
	Email        string    `json:"email" gorm:"column:email;not null" validate:"required,email"`
	Address      Address   `json:"address" gorm:"column:address;not null" validate:"required"`
	CustomerID   uint      `json:"customer" gorm:"column:customerId"`
}
