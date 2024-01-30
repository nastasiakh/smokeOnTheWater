package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type OrderProductRepository struct {
	db *gorm.DB
}

func NewOrderProductRepository(db *gorm.DB) *OrderProductRepository {
	return &OrderProductRepository{db: db}
}

func (repo *OrderProductRepository) Create(body *models.OrderProduct) error {
	if err := validation.ValidateStruct(*body); err != nil {
		return err
	}
	if err := repo.db.Create(body).Error; err != nil {
		return err
	}
	return nil
}

func (repo *OrderProductRepository) FindAll() ([]models.OrderProduct, error) {
	var orderProducts []models.OrderProduct
	return orderProducts, repo.db.Find(&orderProducts).Error
}

func (repo *OrderProductRepository) FindByOrderID(orderId uint) ([]*models.OrderProduct, error) {
	var orderProducts []*models.OrderProduct
	if err := repo.db.Where("order_id = ?", orderId).Find(&orderProducts).Error; err != nil {
		return nil, err
	}
	return orderProducts, nil
}

func (repo *OrderProductRepository) Update(orderID uint, body models.OrderProduct) (*models.OrderProduct, error) {
	if err := validation.ValidateStruct(&body); err != nil {
		return nil, err
	}

	existingProducts, err := repo.FindByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	var index int
	var found bool
	for i, product := range existingProducts {
		if product.ID == body.ID {
			index = i
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("product with ID %d not found in order %d", body.ID, orderID)
	}

	if err := repo.db.Model(&existingProducts[index]).Updates(&body).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func (repo *OrderProductRepository) DeleteByOrderId(orderId uint) error {
	existingProducts, err := repo.FindByOrderID(orderId)
	if err != nil {
		return err
	}

	for _, product := range existingProducts {
		if err := repo.db.Delete(&product).Error; err != nil {
			return err
		}
	}

	return nil
}
