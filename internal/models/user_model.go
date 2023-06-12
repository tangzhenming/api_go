package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string
	Email            string `gorm:"unique"`
	VerificationCode string
	Token            string
}

type ResponseUser struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
}
