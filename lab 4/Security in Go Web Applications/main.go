package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")
var csrfAuthKey = []byte("32-byte-long-auth-key")

var users = map[string]string{
	"admin": hashPassword("password"),
	"user":  hashPassword("password"),
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	return string(hashed)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(username, role string) (string, error) {
	claims := &CustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func getCSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	w.Header().Set("X-CSRF-Token", token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"csrf_token": token})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	storedHash, exists := users[creds.Username]
	if !exists || !checkPasswordHash(creds.Password, storedHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	role := "user"
	if creds.Username == "admin" {
		role = "admin"
	}
	token, err := generateToken(creds.Username, role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func main() {
	router := mux.NewRouter()

	csrfMiddleware := csrf.Protect(
		csrfAuthKey,
		csrf.Secure(false),
	)

	router.HandleFunc("/login", getCSRFToken).Methods("GET")

	router.HandleFunc("/login", loginHandler).Methods("POST")

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", csrfMiddleware(router)))
}
