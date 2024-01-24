package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	return categories, repo.db.Find(&categories).Error
}

func (repo *CategoryRepository) FindById(id uint) (models.Category, error) {
	var category models.Category
	if err := repo.db.First(&category, id).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (repo *CategoryRepository) Create(category *models.Category) (*models.Category, error) {
	if err := validation.ValidateStruct(*category); err != nil {
		return nil, err
	}

	if err := repo.db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (repo *CategoryRepository) Update(id uint, categoryBody *models.Category) (models.Category, error) {
	if err := validation.ValidateStruct(*categoryBody); err != nil {
		return models.Category{}, err
	}

	existingCategory, err := repo.FindById(id)
	if err != nil {
		return models.Category{}, nil
	}

	if err := repo.db.First(&existingCategory, id).Error; err != nil {
		return models.Category{}, err
	}

	if err := repo.db.Model(&existingCategory).Update(categoryBody).Error; err != nil {
		return models.Category{}, err
	}

	return existingCategory, nil
}

func (repo *CategoryRepository) DeleteOne(id uint) error {
	existingCategory, err := repo.FindById(id)
	if err != nil {
		return err
	}

	if err := repo.db.Delete(&existingCategory).Error; err != nil {
		return err
	}
	return nil
}
