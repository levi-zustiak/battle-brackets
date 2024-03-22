package models

import (
	"time"
)

type PrimaryKey struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type Timestamps struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

type User struct {
	PrimaryKey `gorm:"embedded"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"-"`
	Timestamps `gorm:"embedded"`
}
