package db

import (
	"asssignment2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "127.0.0.1"
	user     = "postgres"
	password = "postgres"
	dbPort   = 5432
	dbName   = "ecommerce_assignment2"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
