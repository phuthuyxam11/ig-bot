package db

import (
	"gorm.io/gorm"
	usermodel "igbot.com/authentication/modules/auth/model"
)

func Migrate(dbOrm *gorm.DB) error {
	// register model
	users := usermodel.UsersModel{}

	err := dbOrm.AutoMigrate(&users)
	return err
}
