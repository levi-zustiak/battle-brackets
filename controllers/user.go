package controllers

import (
	"battle-brackets/interfaces"
	"battle-brackets/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {}
func (c *UserController) Show(w http.ResponseWriter, r *http.Request)  {}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var input interfaces.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err = validate.Struct(input)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	user, err := c.service.Create(&input)
	if err != nil {
		http.Error(w, fmt.Sprintln("Error creating user"), http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {}
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {}
