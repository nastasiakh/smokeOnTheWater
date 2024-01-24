package services

import (
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type CategoryService struct {
	categoryRepository *repositories.CategoryRepository
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (service *CategoryService) GetAll() ([]models.Category, error) {
	categories, err := service.categoryRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (service *CategoryService) GetById(id uint) (models.Category, error) {
	category, err := service.categoryRepository.FindById(id)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (service *CategoryService) Create(category *models.Category) (*models.Category, error) {
	if err := validation.ValidateStruct(*category); err != nil {
		return nil, err
	}

	createdCategory, err := service.categoryRepository.Create(category)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (service *CategoryService) Update(id uint, body *models.Category) (models.Category, error) {
	if err := validation.ValidateStruct(body); err != nil {
		return models.Category{}, err
	}

	category, err := service.categoryRepository.Update(id, body)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (service *CategoryService) Delete(id uint) error {
	if err := service.categoryRepository.DeleteOne(id); err != nil {
		return err
	}
	return nil
}
