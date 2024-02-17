package main

import (
	"fmt"
	"github.com/Dubbril/my-gin-project/app/config"
	"github.com/Dubbril/my-gin-project/app/controllers"
	"github.com/Dubbril/my-gin-project/app/database"
	"github.com/Dubbril/my-gin-project/app/exception"
	"github.com/Dubbril/my-gin-project/app/middleware"
	"github.com/Dubbril/my-gin-project/app/repositories"
	"github.com/Dubbril/my-gin-project/app/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	// Init config
	config.GetConfig()

	// Init pattern logger
	middleware.InitLogger()

	// Initialize the database
	db := database.GetDB()

	// Set Gin to release mode to disable debug output
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Use the logger middleware
	r.Use(middleware.LogHandler())
	r.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Initialize repository and service
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Initialize controller
	userController := controllers.NewUserController(userService)

	userApi := r.Group("/api/v1/users")
	{
		userApi.POST("", userController.Create)
		userApi.GET("", userController.GetAll)
		userApi.GET("/:id", userController.GetByID)
		userApi.PUT("/:id", userController.Update)
		userApi.DELETE("/:id", userController.Delete)
	}

	externalApi := r.Group("api/v1/external")
	{
		externalApi.GET("", userController.CallExternal)
	}

	// Run the application
	port := 8080
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot run application")
	}
}
