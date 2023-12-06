package models

type RolePermission struct {
	ID           uint `json:"id" gorm:"private_key"`
	RoleID       uint `json:"roleId"`
	PermissionID uint `json:"permissionId"`
}
