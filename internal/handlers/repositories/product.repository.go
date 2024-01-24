package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	return products, repo.db.Preload("Categories").Find(&products).Error
}

func (repo *ProductRepository) FindById(id uint) (*models.Product, error) {
	var product models.Product
	if err := repo.db.Preload("Categories").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	if err := validation.ValidateStruct(*product); err != nil {
		return nil, err
	}

	tx := repo.db.Begin()

	if err := repo.db.Create(product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, role := range product.Categories {
		if err := tx.Model(&product).Association("Categories").Append(role).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return product, nil
}

func (repo *ProductRepository) Update(id uint, productBody *models.Product) (*models.Product, error) {
	tx := repo.db.Begin()

	if err := validation.ValidateStruct(*productBody); err != nil {
		return nil, err
	}

	existingProduct, err := repo.FindById(id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existingProduct).Updates(productBody).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existingProduct).Association("Categories").Replace(productBody.Categories).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return existingProduct, nil
}

func (repo *ProductRepository) DeleteById(id uint) error {
	tx := repo.db.Begin()

	var existingProduct models.Product
	if err := tx.First(&existingProduct, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&existingProduct).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
