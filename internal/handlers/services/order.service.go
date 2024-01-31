package services

import (
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
	if err := validation.ValidateStruct(*orderBody); err != nil {
		return nil, err
	}

	createdOrder, err := service.orderRepo.Create(orderBody)
	if err != nil {
		return nil, err
	}
	for _, op := range orderProducts {
		op.OrderID = createdOrder.ID
		if err := service.orderProductRepo.Create(op); err != nil {
			return nil, err
		}
		if err := service.quantityCalculatorService.CalculateQuantity(op.ProductID, int(op.Quantity)); err != nil {
			return nil, err
		}
	}
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
	if err := validation.ValidateStruct(updatedOrder); err != nil {
		return models.OrderWithProducts{}, err
	}

	_, err := service.orderRepo.Update(orderID, updatedOrder.Order)
	if err != nil {
		return models.OrderWithProducts{}, err
	}

	existingProducts, err := service.orderProductRepo.FindByOrderID(orderID)
	if err != nil {
		return models.OrderWithProducts{}, err
	}

	for _, updatedProduct := range updatedOrder.OrderProducts {
		var existingProduct *models.OrderProduct
		for _, existing := range existingProducts {

			if existing.ProductID == updatedProduct.ProductID {
				existingProduct = existing
				break
			}
		}

		if existingProduct != nil {
			_, err := service.orderProductRepo.Update(updatedProduct.ID, *updatedProduct)
			if err != nil {
				return models.OrderWithProducts{}, err
			}

			diff := int(updatedProduct.Quantity) - int(existingProduct.Quantity)
			if diff != 0 {
				if err := service.quantityCalculatorService.CalculateQuantity(updatedProduct.ProductID, diff); err != nil {
					return models.OrderWithProducts{}, err
				}
			}
		} else {
			if err := service.orderProductRepo.Create(updatedProduct); err != nil {
				return models.OrderWithProducts{}, err
			}
		}
	}

	return updatedOrder, nil
}

func (service *OrderService) Delete(orderId uint) error {
	if err := service.orderRepo.DeleteById(orderId); err != nil {
		return err
	}

	if err := service.orderProductRepo.DeleteByOrderId(orderId); err != nil {
		return err
	}

	return nil
}
