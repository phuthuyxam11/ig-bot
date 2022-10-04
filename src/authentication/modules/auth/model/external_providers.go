package usermodel

import (
	"igbot.com/authentication/utils"
)

type ExternalProviders struct {
	utils.SQLModel `json:",inline"`
	ProviderName   string `json:"provider_name" gorm:"not null;column:provider_name"`
	WSEEndpoint    string `json:"wse_endpoint" gorm:"not null;column:wse_endpoint"`
}
