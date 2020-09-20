package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin" gorm:"default:false"`
}
