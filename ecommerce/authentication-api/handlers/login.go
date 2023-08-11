package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"ecommerce_api.com/auth/authentication-api/models"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var cmd models.LoginUserCommand
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	storedPassword, exists := registeredUsers[cmd.Username]
	if !exists || storedPassword != cmd.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(cmd.Username)
	if err != nil {
		http.Error(w, "Token generation error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Login successful", "token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
