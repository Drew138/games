package models

import "time"

// Queries - Information pertaining searches done by each user
type Queries struct {
	ID        int `gorm:"primaryKey"`
	Query     string
	Result    string
	CreatedAt time.Time
	UserInfo  User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
