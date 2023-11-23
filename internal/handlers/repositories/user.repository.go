package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	return users, repo.db.Find(&users).Error
}

func (repo *UserRepository) FindById(userId uint) (models.User, error) {
	var user models.User
	if err := repo.db.First(&user, userId).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) Update(userId uint, userBody models.User) (models.User, error) {
	var existingUser models.User

	if err := repo.db.First(&existingUser, userId).Error; err != nil {
		return models.User{}, err
	}

	if err := repo.db.Model(&existingUser).Update(userBody).Error; err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

func (repo *UserRepository) DeleteById(userId uint) error {
	var existingUser models.User

	if err := repo.db.First(&existingUser, userId).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&existingUser).Error; err != nil {
		return err
	}

	return nil
}
