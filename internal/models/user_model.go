package models

import "time"

type User struct {
	UserID     int       `gorm:"primary_key;autoIncrement" json:"userID"`
	UserName   string    `json:"userName"`
	CreateTime time.Time `json:"createTime"`
}
