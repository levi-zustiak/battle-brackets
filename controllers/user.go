package controllers

import (
	"battle-brackets/models"
	"battle-brackets/repos"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type CreateUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserController struct {
	repo repos.UserRepo
}

func NewUserController(repo repos.UserRepo) *UserController {
	return &UserController{repo: repo}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	// return c.repo.GetAll()
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateUserInput
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

	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}

	c.repo.Create(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {}
