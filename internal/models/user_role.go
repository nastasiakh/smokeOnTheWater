package models

type UserRole struct {
	ID     uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"userId"`
	RoleID uint `json:"roleId"`
}
