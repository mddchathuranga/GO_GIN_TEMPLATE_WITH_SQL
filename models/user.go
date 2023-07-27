package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `gorm:"not null"`
	Address  string `gorm:"not null"`
	Mobile   string `gorm:"not null;unique"`
	Age      int    `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	// Define the association with posts
	Posts []Post `gorm:"foreignKey:UserID"`
}
