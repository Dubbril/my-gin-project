package services

import (
	"encoding/json"
	"github.com/Dubbril/my-gin-project/app/config"
	"github.com/Dubbril/my-gin-project/app/domain"
	"github.com/Dubbril/my-gin-project/app/helper"
	"github.com/Dubbril/my-gin-project/app/middleware"
	"github.com/Dubbril/my-gin-project/app/models"
	"github.com/Dubbril/my-gin-project/app/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID uuid.UUID) error
}

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepository.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepository.GetAllUsers()
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	return s.UserRepository.GetUserByID(userID)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(userID uuid.UUID) error {
	return s.UserRepository.DeleteUser(userID)
}

func (s *UserService) CallExternal() domain.ExternalResponse {
	getConfig := config.GetConfig()
	newResty := middleware.NewRestClient(resty.New())
	resp, err := newResty.Client.R().Get(getConfig.ExternalUrl)
	if err != nil {
		panic(err)
	}

	if resp.IsError() {
		msgError, err := helper.ConvByteArrayToJsonOrArray(resp.Body())
		if err != nil {
			panic(string(resp.Body()))
		}
		panic(msgError)
	}

	// Bind data to response struct
	var responseData domain.ExternalResponse
	err = json.Unmarshal(resp.Body(), &responseData)
	if err != nil {
		panic(err)
	}

	// Use the validate to check for required fields
	validate := validator.New()
	err = validate.Struct(responseData)
	if err != nil {
		panic(err)
	}

	return responseData
}
