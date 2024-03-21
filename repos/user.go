package repos

import (
	"battle-brackets/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db}
}

func (r UserRepo) Create(u *models.User) *models.User {
	if result := r.DB.Create(&u); result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(u)

	return u
}

func (u UserRepo) GetOne() {}
