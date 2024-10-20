package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"midterm/controllers"
	"midterm/database"
	"net/http"
)

func main() {
	database.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/items", controllers.GetItems).Methods("GET")
	r.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(r)))
}
