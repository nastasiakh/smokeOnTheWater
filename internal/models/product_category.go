package models

type ProductCategory struct {
	ID         uint `json:"id" gorm:"primary_key"`
	ProductID  uint `json:"productId"`
	CategoryID uint `json:"categoryId"`
}
