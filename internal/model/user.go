package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"size:255,not null"`
	Email    string `gorm:"size:100,unique"`
}
