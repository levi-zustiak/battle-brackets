package controllers

import (
	"battle-brackets/models"
	"battle-brackets/repos"
	"encoding/json"
	"net/http"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
	json.NewDecoder(r.Body).Decode(&input)

	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}

	c.repo.Create(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {}
