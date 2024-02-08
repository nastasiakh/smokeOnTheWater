package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
	"time"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (repo *OrderRepository) Create(tx *gorm.DB, order *models.Order) (*models.Order, error) {

	if err := validation.ValidateStruct(*order); err != nil {
		return nil, err
	}

	if err := tx.Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (repo *OrderRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order
	return orders, repo.Db.Find(&orders).Error
}

func (repo *OrderRepository) FindById(orderId uint) (models.Order, error) {
	var order models.Order
	if err := repo.Db.First(&order, orderId).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (repo *OrderRepository) Update(tx *gorm.DB, orderId uint, orderBody models.Order) (models.Order, error) {

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

	if err := tx.Save(&existingOrder).Error; err != nil {
		return models.Order{}, err
	}
	return existingOrder, nil
}

func (repo *OrderRepository) DeleteById(tx *gorm.DB, id uint) error {
	var existingOrder models.Order

	if err := tx.First(&existingOrder, id).Error; err != nil {
		return err
	}

	if err := tx.Delete(&existingOrder).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
