package models

import (
	"gorm.io/gorm"
)

type Users struct {
	// utilizado para transformar a struct em tabel no DB gorm.Model
	gorm.Model
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

func NewUser() *Users {
	users := Users{}
	return &users
}
