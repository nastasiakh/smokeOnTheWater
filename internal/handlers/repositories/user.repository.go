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
	return users, repo.db.Find(&users).Error
}

func (repo *UserRepository) FindById(id uint) (models.User, error) {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) error {
	if err := validation.ValidateStruct(*user); err != nil {
		return err
	}

	if len(user.Roles) == 0 {
		role := models.Role{}
		if err := repo.db.Where(models.Role{Name: "client"}).FirstOrCreate(&role).Error; err != nil {
			return err
		}
		user.Roles = append(user.Roles, role)
	}

	return repo.db.Create(user).Error
}

func (repo *UserRepository) Update(id uint, userBody *models.User) (models.User, error) {
	var existingUser models.User
	if err := validation.ValidateStruct(*userBody); err != nil {
		return models.User{}, err
	}

	if err := repo.db.First(&existingUser, id).Error; err != nil {
		return models.User{}, err
	}

	if err := repo.db.Model(&existingUser).Update(userBody).Error; err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

func (repo *UserRepository) DeleteById(id uint) error {
	var existingUser models.User

	if err := repo.db.First(&existingUser, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&existingUser).Error; err != nil {
		return err
	}

	return nil
}
