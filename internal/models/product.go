package models

type Product struct {
	ID              uint       `json:"id" gorm:"primary key"`
	Title           string     `json:"title" gorm:"column:title;not null" validate:"required"`
	Description     string     `json:"description" gorm:"column:description;not null" validate:"required"`
	MetaTitle       string     `json:"metaTitle" gorm:"column:meta_title;not null" validate:"required"`
	MetaDescription string     `json:"metaDescription" gorm:"column:meta_description;not null" validate:"required"`
	Sku             string     `json:"sku" gorm:"column:sku;unique;not null" validate:"required"`
	Quantity        uint       `json:"quantity" gorm:"column:quantity;not null" validate:"required"`
	Images          string     `json:"images" gorm:"column:images;not null;" validate:"required"`
	Categories      []Category `gorm:"many2many:product_category;" json:"categories"`
}
