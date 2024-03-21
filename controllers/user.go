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

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	c.repo.Create(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}
