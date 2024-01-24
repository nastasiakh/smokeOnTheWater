package services

import (
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (service *ProductService) GetAll() ([]models.Product, error) {
	products, err := service.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service *ProductService) GetById(id uint) (*models.Product, error) {
	product, err := service.productRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) Create(product *models.Product) (*models.Product, error) {

	if err := validation.ValidateStruct(*product); err != nil {
		return nil, err
	}

	createdProduct, err := service.productRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return createdProduct, nil
}

func (service *ProductService) Update(id uint, body *models.Product) (*models.Product, error) {
	if err := validation.ValidateStruct(body); err != nil {
		return nil, err
	}
	product, err := service.productRepository.Update(id, body)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (service *ProductService) Delete(id uint) error {
	if err := service.productRepository.DeleteById(id); err != nil {
		return err
	}
	return nil
}
