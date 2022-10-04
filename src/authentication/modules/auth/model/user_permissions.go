package usermodel

import (
	"igbot.com/authentication/utils"
)

type UserPermissions struct {
	utils.SQLModel  `json:",inline"`
	RoleName        string `json:"provider_name" gorm:"not null;column:role_name"`
	RoleDescription string `json:"wse_endpoint" gorm:"column:role_description"`
}
