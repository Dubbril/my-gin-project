package main

import (
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/controllers"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/database"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/repositories"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	// Initialize the database
	database.Init()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(database.GetDB())

	// Create Gin router
	r := gin.Default()

	// Initialize repository and service
	userRepo := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepo)

	// Initialize controller
	userController := controllers.NewUserController(userService)

	// Define routes
	r.POST(`/users`, userController.Create)
	r.GET(`/users`, userController.GetAll)
	r.GET(`/users/:id`, userController.GetByID)
	r.PUT(`/users/:id`, userController.Update)
	r.DELETE(`/users/:id`, userController.Delete)

	// Run the application
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
