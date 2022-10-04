package db

import (
	"gorm.io/gorm"
	usermodel "igbot.com/authentication/modules/auth/model"
)

func Migrate(dbOrm *gorm.DB) error {
	// register model
	userLogin := usermodel.UserLoginData{}

	err := dbOrm.AutoMigrate(&userLogin)
	return err
}
