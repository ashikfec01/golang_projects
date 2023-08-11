package main

import (
	"net/http"

	"ecommerce_api.com/auth/authentication-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUserHandler).Methods("POST")

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
