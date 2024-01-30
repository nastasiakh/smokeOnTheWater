package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
	"time"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (repo *OrderRepository) Create(order *models.Order) (*models.Order, error) {
	if err := validation.ValidateStruct(*order); err != nil {
		return nil, err
	}

	if err := repo.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *OrderRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order
	return orders, repo.db.Find(&orders).Error
}

func (repo *OrderRepository) FindById(orderId uint) (models.Order, error) {
	var order models.Order
	if err := repo.db.First(&order, orderId).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (repo *OrderRepository) Update(orderId uint, orderBody models.Order) (models.Order, error) {
	if err := validation.ValidateStruct(&orderBody); err != nil {
		return models.Order{}, err
	}

	existingOrder, err := repo.FindById(orderId)
	if err != nil {
		return models.Order{}, err
	}
	existingOrder.TotalAmount = orderBody.TotalAmount
	existingOrder.DateModified = time.Now()
	existingOrder.Status = orderBody.Status
	existingOrder.FirstName = orderBody.FirstName
	existingOrder.LastName = orderBody.LastName
	existingOrder.Phone = orderBody.Phone
	existingOrder.Email = orderBody.Email
	existingOrder.Address = orderBody.Address
	existingOrder.CustomerID = orderBody.CustomerID

	if err := repo.db.Save(&existingOrder).Error; err != nil {
		return models.Order{}, err
	}

	return existingOrder, err

}

func (repo *OrderRepository) DeleteById(id uint) error {
	var existingOrder models.Order

	if err := repo.db.First(&existingOrder, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&existingOrder).Error; err != nil {
		return err
	}
	return nil
}
