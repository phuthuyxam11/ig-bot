package usermodel

import (
	"igbot.com/authentication/utils"
)

type Permissions struct {
	utils.SQLModel        `json:",inline"`
	PermissionName        string `json:"provider_name" gorm:"not null;column:role_name"`
	PermissionDescription string `json:"wse_endpoint" gorm:"column:role_description"`
}
