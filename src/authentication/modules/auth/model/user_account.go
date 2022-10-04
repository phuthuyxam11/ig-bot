package usermodel

import (
	"gorm.io/gorm"
	"time"
)

type UserAccount struct {
	gorm.Model
	FirstName   string `gorm:"not null"`
	LastName    string
	Gender      int `gorm:"size:1"`
	DateOfBirth time.Time
	RoleId      int `gorm:"not null"`
}
