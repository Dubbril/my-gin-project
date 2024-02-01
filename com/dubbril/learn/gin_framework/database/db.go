package database

import (
	"fmt"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/config"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Init initializes the database connection
func Init() {
	cfg := config.NewDatabaseConfig()

	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s search_path=%s",
		cfg.Host, cfg.User, cfg.DBName, cfg.SSLMode, cfg.Password, cfg.SearchPath))
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connected to the database")

	// Enable Gorm's auto migration feature to automatically create tables based on the struct models
	db.AutoMigrate(&models.User{}) // Uncomment this line when you have models to migrate
}

// GetDB returns the Gorm DB instance
func GetDB() *gorm.DB {
	return db
}
