package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint   // Foreign key to associate the post with a user
	User    User   `gorm:"foreignKey:UserID"`
}
