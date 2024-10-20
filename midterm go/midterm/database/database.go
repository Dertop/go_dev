package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"midterm/models"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123 dbname=midterm port=5432 sslmode=disable TimeZone=Europe/Moscow"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connection established")

	err = DB.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
