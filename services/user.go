package services

import (
	"battle-brackets/interfaces"
	"battle-brackets/models"
	"battle-brackets/repos"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repos.UserRepo
}

func NewUserService(repo *repos.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) Create(i *interfaces.CreateUserInput) (models.User, error) {
	// Hash the password here
	hash, _ := HashPassword(i.Password)

	fmt.Println(hash)
	u := models.User{Name: i.Name, Email: i.Email, Password: hash}

	s.repo.Create(&u)

	return u, nil
}
