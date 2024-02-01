// services/user_service.go
package services

import (
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/models"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

// CreateUser Create a new user
func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepository.CreateUser(user)
}

// GetAllUsers Get all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepository.GetAllUsers()
}

// GetUserByID Get a user by ID
func (s *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	return s.UserRepository.GetUserByID(userID)
}

// UpdateUser Update a user
func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepository.UpdateUser(user)
}

// DeleteUser Delete a user by ID
func (s *UserService) DeleteUser(userID uuid.UUID) error {
	return s.UserRepository.DeleteUser(userID)
}
