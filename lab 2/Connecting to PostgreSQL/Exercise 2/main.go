package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func DBConnect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func applyMigrations(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func addNewUser(db *gorm.DB, name, email, password string) {
	newUser := User{Name: name, Email: email, Password: password}
	db.Create(&newUser)
}

func fetchUsers(db *gorm.DB) ([]User, error) {
	var allUsers []User
	result := db.Find(&allUsers)
	return allUsers, result.Error
}

func main() {
	db, err := DBConnect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	applyMigrations(db)

	addNewUser(db, "Anuar Daniyar", "dertop125@gmail.com", "12354")
	addNewUser(db, "Dosimzhan Kogabaev", "chel@gmail.com", "2222")

	users, err := fetchUsers(db)
	if err != nil {
		log.Fatal("Error fetching users:", err)
	}

	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
