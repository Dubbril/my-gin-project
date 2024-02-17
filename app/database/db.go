package database

import (
	"github.com/Dubbril/my-gin-project/app/config"
	"github.com/Dubbril/my-gin-project/app/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {

	getConfig := config.GetConfig()
	postgresConn := postgres.Open(getConfig.PostgresConnection)
	db, err := gorm.Open(postgresConn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	log.Info().Msg("Connected to the database success !!!")

	// Auto create table
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to auto create table")
	}

	return db
}
