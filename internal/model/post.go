package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(255);not null"`
	Content  string    `gorm:"type:text"`
	UserID   uint      `gorm:"foreignKey:UserID"`
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}
