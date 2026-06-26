package controller

import (
	"encoding/json"
	"fmt"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/services"
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
	// err = c.service.SignUp(user)
	// if err != nil {
	// 	fmt.Println("SIGNUP ERROR:", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	err = c.service.SignUp(user)
	if err != nil {
		fmt.Println("SIGNUP ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User Created")
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
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, "Login Successful")
}
