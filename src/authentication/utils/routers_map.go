package utils

import "igbot.com/authentication/configs"

func GetRoutersMap() map[string]string {
	config := configs.LoadConfig()
	return map[string]string{
		"auth_signup":         config.DOMAIN + "/auth/signup",
		"auth_login":          config.DOMAIN + "/auth/login",
		"auth_profile":        config.DOMAIN + "/auth/profile",
		"auth_verify_acc":     config.DOMAIN + "/auth/verify-acc",
		"auth_password_reset": config.DOMAIN + "/auth/password-reset",
	}
}
