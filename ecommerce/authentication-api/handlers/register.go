package handlers

import (
	"encoding/json"
	"net/http"

	"ecommerce_api.com/auth/authentication-api/models"
)

// Mock data store for registered users
var registeredUsers = map[string]string{} // username -> password

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var cmd models.RegisterUserCommand
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if _, exists := registeredUsers[cmd.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	registeredUsers[cmd.Username] = cmd.Password

	response := map[string]string{"message": "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
