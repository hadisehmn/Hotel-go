package controller

import (
	"encoding/json"
	"fmt"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/services"
	"log"
	"net/http"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}
func (c *UserController) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.service.SignUp(user)
	if err != nil {
		log.Printf("SignUp failed: %v", err)

		if err.Error() == "user already exists" {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func (c *UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = c.service.SignIn(user)
	if err != nil {
		log.Printf("SignIn failed: %v", err)

		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err.Error() == "wrong password" {
			http.Error(w, "Wrong password", http.StatusUnauthorized)
			return
		}

		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, "Login successful")
}
