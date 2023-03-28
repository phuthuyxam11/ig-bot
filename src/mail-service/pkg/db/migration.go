package db

import (
	"gorm.io/gorm"
)

func Migrate(dbOrm *gorm.DB) error {
	// register model
	//userLogin := usermodel.UserLoginData{}
	//externalProviders := usermodel.ExternalProviders{}
	//grantedPermissions := usermodel.GrantedPermissions{}
	//passwordReset := usermodel.PasswordResetsModel{}
	//permissions := usermodel.Permissions{}
	//userAcc := usermodel.UserAccount{}
	//userLoginExternal := usermodel.UserLoginDataExternal{}
	//userPermissions := usermodel.UserPermissions{}
	//userRoles := usermodel.UserRoles{}
	//
	//err := dbOrm.AutoMigrate(&userLogin,
	//	&externalProviders,
	//	&grantedPermissions,
	//	&passwordReset,
	//	&permissions,
	//	&userAcc,
	//	&userLoginExternal,
	//	&userPermissions,
	//	&userRoles)
	return nil
}
