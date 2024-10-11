package main

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func fetchAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Database connection failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var userList []User
	result := db.Find(&userList)
	if result.Error != nil {
		http.Error(w, "Error fetching users: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userList)
}

func insertNewUser(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Database connection failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var newUser User
	if decodeErr := json.NewDecoder(r.Body).Decode(&newUser); decodeErr != nil {
		http.Error(w, "Invalid input: "+decodeErr.Error(), http.StatusBadRequest)
		return
	}

	if createErr := db.Create(&newUser); createErr.Error != nil {
		http.Error(w, "Failed to create user: "+createErr.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func modifyUser(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Database connection failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userID := mux.Vars(r)["id"]

	var updatedUser User
	if decodeErr := json.NewDecoder(r.Body).Decode(&updatedUser); decodeErr != nil {
		http.Error(w, "Invalid input: "+decodeErr.Error(), http.StatusBadRequest)
		return
	}

	result := db.Model(&User{}).Where("id = ?", userID).Updates(updatedUser)
	if result.Error != nil {
		http.Error(w, "Error updating user: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Database connection failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userID := mux.Vars(r)["id"]

	result := db.Delete(&User{}, userID)
	if result.Error != nil {
		http.Error(w, "Error deleting user: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", fetchAllUsers).Methods("GET")
	router.HandleFunc("/api/user", insertNewUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", modifyUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", removeUser).Methods("DELETE")

	server := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server is running on port 8000...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
}
