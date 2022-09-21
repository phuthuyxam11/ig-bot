package usermodel

import "gorm.io/gorm"

type PasswordResetsModel struct {
	gorm.Model
	Email string `gorm:"not null"`
	Token string `gorm:"not null"`
}
