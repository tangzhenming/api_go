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
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
}
