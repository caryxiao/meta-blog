package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	UserID  uint   `gorm:"foreignKey:UserID"`
	PostID  uint   `gorm:"foreignKey:PostID"`
}
