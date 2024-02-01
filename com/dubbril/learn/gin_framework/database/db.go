package database

import (
	"fmt"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Init initializes the database connection
func Init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=bit@1234 search_path=test")
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
