package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"midterm/database"
	"midterm/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	result := database.DB.Find(&items)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&item)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.Item
	result := database.DB.First(&item, id)
	if result.Error != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result = database.DB.Save(&item)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.Item
	result := database.DB.First(&item, id)
	if result.Error != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	database.DB.Delete(&item)
	fmt.Fprintf(w, "Item successfully deleted")
}
