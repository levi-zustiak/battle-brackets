package repos

import (
	"battle-brackets/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepo struct {
	pg *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{pg: db}
}

func (r UserRepo) Create(u *models.User) *models.User {
	if result := r.pg.Create(&u); result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(u)

	return u
}

func (u UserRepo) GetOne() {}

func (u UserRepo) GetAll() {}

func (u *UserRepo) Update() {}

func (u *UserRepo) Delete() {}
