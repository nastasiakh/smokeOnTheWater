package services

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type OrderService struct {
	orderRepo                 *repositories.OrderRepository
	orderProductRepo          *repositories.OrderProductRepository
	quantityCalculatorService *QuantityCalculatorService
}

func NewOrderService(
	orderRepository *repositories.OrderRepository,
	orderProductRepository *repositories.OrderProductRepository,
	quantityCalculatorService *QuantityCalculatorService) *OrderService {
	return &OrderService{
		orderRepo:                 orderRepository,
		orderProductRepo:          orderProductRepository,
		quantityCalculatorService: quantityCalculatorService,
	}
}

func (service *OrderService) Create(orderBody *models.Order, orderProducts []*models.OrderProduct) (*models.Order, error) {
	tx := service.orderRepo.Db.Begin()

	if err := validation.ValidateStruct(*orderBody); err != nil {
		tx.Rollback()
		return nil, err
	}

	createdOrder, err := service.orderRepo.Create(tx, orderBody)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, op := range orderProducts {
		op.OrderID = createdOrder.ID
		if err := service.orderProductRepo.Create(tx, op); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := service.quantityCalculatorService.CalculateQuantity(op.ProductID, int(op.Quantity)); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return createdOrder, err
}

func (service *OrderService) GetAll() ([]models.OrderWithProducts, error) {
	var ordersWithProducts []models.OrderWithProducts

	orders, err := service.orderRepo.FindAll()
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		products, err := service.orderProductRepo.FindByOrderID(order.ID)
		if err != nil {
			return nil, err
		}

		orderWithProducts := models.OrderWithProducts{
			Order:         order,
			OrderProducts: products,
		}

		ordersWithProducts = append(ordersWithProducts, orderWithProducts)
	}

	return ordersWithProducts, nil
}

func (service *OrderService) GetById(id uint) (models.OrderWithProducts, error) {
	var orderWithProduct models.OrderWithProducts

	order, err := service.orderRepo.FindById(id)
	if err != nil {
		return models.OrderWithProducts{}, err
	}

	products, err := service.orderProductRepo.FindByOrderID(order.ID)
	if err != nil {
		return models.OrderWithProducts{}, err
	}

	orderWithProduct = models.OrderWithProducts{
		Order:         order,
		OrderProducts: products,
	}

	return orderWithProduct, nil
}

func (service *OrderService) Update(orderID uint, updatedOrder models.OrderWithProducts) (models.OrderWithProducts, error) {
	tx := service.orderRepo.Db.Begin()

	if err := validation.ValidateStruct(updatedOrder); err != nil {
		tx.Rollback()
		return models.OrderWithProducts{}, err
	}

	_, err := service.orderRepo.Update(tx, orderID, updatedOrder.Order)
	if err != nil {
		tx.Rollback()
		return models.OrderWithProducts{}, err
	}

	existingProducts, err := service.orderProductRepo.FindByOrderID(orderID)
	if err != nil {
		tx.Rollback()
		return models.OrderWithProducts{}, err
	}

	if err := service.updateOrderProducts(tx, orderID, existingProducts, updatedOrder.OrderProducts); err != nil {
		tx.Rollback()
		return models.OrderWithProducts{}, err
	}

	tx.Commit()
	return updatedOrder, nil
}

func (service *OrderService) updateOrderProducts(tx *gorm.DB, orderId uint, existingProducts, updatedProducts []*models.OrderProduct) error {
	productMap := buildProductMap(existingProducts)

	for _, updatedProduct := range updatedProducts {
		prevQuantity, found := productMap[updatedProduct.ProductID]
		if found {
			_, err := service.orderProductRepo.Update(tx, updatedProduct.OrderID, *updatedProduct)
			if err != nil {
				return err
			}

			diff := calculateQuantityDifference(updatedProduct.Quantity, prevQuantity)
			if diff != 0 {
				if err := service.quantityCalculatorService.CalculateQuantity(updatedProduct.ProductID, diff); err != nil {
					return err
				}
			}
			delete(productMap, updatedProduct.ProductID)

		} else {
			if err := service.orderProductRepo.Create(tx, updatedProduct); err != nil {
				return err
			}
		}
	}

	if err := service.updateQuantityAndDeleteProducts(tx, orderId, productMap); err != nil {
		return err
	}

	return nil
}

func (service *OrderService) Delete(orderId uint) error {
	tx := service.orderRepo.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	defer tx.Commit()

	if err := service.orderRepo.DeleteById(tx, orderId); err != nil {
		return err
	}

	if err := service.orderProductRepo.DeleteAllByOrderId(tx, orderId); err != nil {
		return err
	}

	return nil
}

func (service *OrderService) updateQuantityAndDeleteProducts(tx *gorm.DB, orderID uint, productMap map[uint]uint) error {
	for productID, quantity := range productMap {
		if err := service.quantityCalculatorService.CalculateQuantity(productID, int(quantity)); err != nil {
			return err
		}
		if err := service.orderProductRepo.DeleteOneByProductId(tx, orderID, productID); err != nil {
			return err
		}
	}
	return nil
}

func buildProductMap(existingProducts []*models.OrderProduct) map[uint]uint {
	productMap := make(map[uint]uint)
	for _, existingProduct := range existingProducts {
		productMap[existingProduct.ProductID] = existingProduct.Quantity
	}
	return productMap
}

func calculateQuantityDifference(newQuantity, oldQuantity uint) int {
	return int(newQuantity) - int(oldQuantity)
}
