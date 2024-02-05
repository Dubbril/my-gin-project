// main.go
package main

import (
	"fmt"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/controllers"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/database"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/middleware"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/repositories"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// Load Config
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Initialize the database
	//database.Init()
	db := database.GetDB()

	// Set Gin to release mode to disable debug output
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	//r := gin.Default()
	r := gin.New()

	// Use the logger middleware
	r.Use(middleware.LogHandler())

	// Initialize repository and service
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Initialize controller
	userController := controllers.NewUserController(userService)

	// Define routes
	r.POST("/users", userController.Create)
	r.GET("/users", userController.GetAll)
	r.GET("/users/:id", userController.GetByID)
	r.PUT("/users/:id", userController.Update)
	r.DELETE("/users/:id", userController.Delete)

	// Run the application
	port := 8080
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatal(err)
	}
}
