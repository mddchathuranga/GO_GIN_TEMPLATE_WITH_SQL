package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"not null;unique"`
	Username string `gorm:"not null"`
	Address  string `gorm:"not null"`
	Mobile   string `gorm:"not null;unique"`
	Age      int    `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
}
