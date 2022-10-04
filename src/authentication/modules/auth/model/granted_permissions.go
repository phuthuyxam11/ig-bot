package usermodel

type GrantedPermissions struct {
	RoleId       int `json:"role_id" gorm:"not null;column:role_id"`
	PermissionId int `json:"permission_id" gorm:"not null;column:permission_id"`
}
