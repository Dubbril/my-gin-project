// repositories/user_repository.go
package repositories

import (
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser Create a new user
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// GetAllUsers Get all users
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID Get a user by ID
func (r *UserRepository) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser Update a user
func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

// DeleteUser Delete a user by ID
func (r *UserRepository) DeleteUser(userID uuid.UUID) error {
	return r.DB.Where("id = ?", userID).Delete(&models.User{}).Error
}
