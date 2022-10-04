package usermodel

type UserRoles struct {
	UserId       int `json:"user_id" gorm:"not null;column:user_id"`
	PermissionId int `json:"permission_id" gorm:"not null;column:permission_id"`
}
