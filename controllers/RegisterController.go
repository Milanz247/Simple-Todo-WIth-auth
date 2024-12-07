package controllers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"Golang/models"
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Check if the Content-Type is application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	var user models.User

	// Decode the JSON request body into the user object
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input, unable to parse JSON", http.StatusBadRequest)
		return
	}

	// Check if Email and Password are populated
	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and Password cannot be empty", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := models.CreateUser(&user); err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
