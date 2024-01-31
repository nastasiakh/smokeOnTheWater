package services

import (
	"smokeOnTheWater/internal/handlers/repositories"
)

type QuantityCalculatorService struct {
	productRepo *repositories.ProductRepository
}

func NewQuantityCalculatorService(productRepo *repositories.ProductRepository) *QuantityCalculatorService {
	return &QuantityCalculatorService{productRepo: productRepo}
}

func (service *QuantityCalculatorService) CalculateQuantity(productId uint, difference int) error {
	product, err := service.productRepo.FindById(productId)
	if err != nil {
		return err
	}

	if difference < 0 {
		product.Quantity += uint(-difference)
	} else {
		product.Quantity -= uint(difference)
	}

	_, err = service.productRepo.Update(product.ID, product)
	if err != nil {
		return err
	}

	return nil
}
