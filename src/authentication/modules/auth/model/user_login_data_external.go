package usermodel

import (
	"igbot.com/authentication/utils"
)

type UserLoginDataExternal struct {
	utils.SQLModel        `json:",inline"`
	ExternalProviderId    int `json:"external_provider_id" gorm:"not null;column:external_provider_id"`
	ExternalProviderToken int `json:"external_provider_token" gorm:"not null;column:external_provider_token"`
}
