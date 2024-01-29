package models

type Category struct {
	ID              uint   `json:"id" gorm:"primary key"`
	Title           string `json:"title" gorm:"unique;not null" validate:"required"`
	Description     string `json:"description" gorm:"not null" validate:"required"`
	MetaTitle       string `json:"metaTitle"  gorm:"unique;not null" validate:"required"`
	MetaDescription string `json:"metaDescription"  gorm:"not null" validate:"required"`
	Image           string `json:"image"  gorm:"not null"`
	ParentID        uint   `json:"parentId" gorm:"foreignKey:ParentID"`
}
