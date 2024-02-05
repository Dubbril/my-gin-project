package database

import (
	"fmt"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/config"
	"github.com/Dubbril/my-gin-project/com/dubbril/learn/gin_framework/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	// Use Viper to bind the configuration to a struct
	var dbConfig config.DbConfig
	if err := viper.Unmarshal(&dbConfig); err != nil {
		panic(fmt.Errorf("Error unmarshaling config: %s \n", err))
	}

	postgresConn := postgres.Open(dbConfig.Postgres.Connection)
	db, err := gorm.Open(postgresConn)
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connected to the database success !!!")

	// Auto create table
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}

	return db
}
