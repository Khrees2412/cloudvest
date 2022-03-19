package database

import (
	"os"

	"cloudvest/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	DB_URI := os.Getenv("DB_URI")
	var err error
	DB, err = gorm.Open(postgres.Open(DB_URI), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	log.Println("Database connected...")

	DB.AutoMigrate(&models.User{}, &models.Claims{}, &models.File{}, &models.Folder{})
}
