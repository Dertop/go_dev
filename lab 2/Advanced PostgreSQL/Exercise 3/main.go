package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	Age  int    `json:"age" gorm:"not null"`
}

func connectGORM() {
	dsn := "host=localhost user=postgres password=123 dbname=postgres2 port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	ageFilter := r.URL.Query().Get("age")
	sortBy := r.URL.Query().Get("sortBy")

	query := DB
	if ageFilter != "" {
		query = query.Where("age = ?", ageFilter)
	}
	if sortBy == "name" {
		query = query.Order("name asc")
	}
	query.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	var existingUser User
	result := DB.Where("name = ?", user.Name).First(&existingUser)
	if result.Error == nil {
		http.Error(w, "User with this name already exists", http.StatusBadRequest)
		return
	}
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var user User
	DB.First(&user, id)
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var user User
	DB.Delete(&user, id)
	json.NewEncoder(w).Encode(fmt.Sprintf("User with ID %d has been deleted", id))
}

func main() {
	connectGORM()
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
