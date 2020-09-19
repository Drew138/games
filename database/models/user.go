package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

// https://stackoverflow.com/questions/23589564/function-for-converting-a-struct-to-map-in-golang
