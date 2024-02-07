package models

type OrderProduct struct {
	ID        uint    `json:"id" gorm:"primary key"`
	OrderID   uint    `json:"orderId" gorm:"column:order_id;foreignKey:OrderID;not null" validate:"required"`
	ProductID uint    `json:"productId" gorm:"column:product_id;not null" validate:"required"`
	Title     string  `json:"title" gorm:"column:title;not null" validate:"required"`
	Quantity  uint    `json:"quantity" gorm:"column:quantity;not null" validate:"required"`
	Price     float64 `json:"price" gorm:"column:price;not null" validate:"required"`
	Sku       string  `json:"sku" gorm:"column:sku;not null" validate:"required"`
}
