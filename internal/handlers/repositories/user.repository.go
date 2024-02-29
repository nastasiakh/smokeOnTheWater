package repositories

import (
	"github.com/jinzhu/gorm"
	"smokeOnTheWater/internal/handlers/validation"
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
	return users, repo.db.Preload("Roles").Preload("Roles.Permissions").Find(&users).Error
}

func (repo *UserRepository) FindById(id uint) (models.User, error) {
	var user models.User
	if err := repo.db.Preload("Roles").Preload("Roles.Permissions").First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Preload("Roles").Preload("Roles.Permissions").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := validation.ValidateStruct(*user); err != nil {
		return nil, err
	}

	tx := repo.db.Begin()

	if err := repo.db.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, role := range user.Roles {
		if err := tx.Model(&user).Association("Roles").Append(role).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return user, nil
}

func (repo *UserRepository) Update(id uint, userBody *models.User) (models.User, error) {
	tx := repo.db.Begin()

	if err := validation.ValidateStruct(*userBody); err != nil {
		return models.User{}, err
	}

	existingUser, err := repo.FindById(id)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	if err := tx.Model(&existingUser).Updates(models.User{
		FirstName: userBody.FirstName,
		LastName:  userBody.LastName,
		Phone:     userBody.Phone,
		Email:     userBody.Email,
	}).Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	if err := tx.Model(&existingUser).Association("Roles").Replace(userBody.Roles).Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	return existingUser, nil
}

func (repo *UserRepository) DeleteById(id uint) error {
	tx := repo.db.Begin()

	var existingUser models.User
	if err := tx.First(&existingUser, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&existingUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
