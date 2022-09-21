package usermodel

import (
	"errors"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component/tokenprovider"
	"igbot.com/authentication/utils"
	"time"
)

const EntityName = "User"

type UsersModel struct {
	utils.SQLModel  `json:",inline"`
	Name            string    `json:"name" gorm:"not null;column:name"`
	Email           string    `json:"email" gorm:"column:email;unique"`
	Salt            string    `json:"-" gorm:"column:salt"`
	EmailVerifiedAt time.Time `json:"-" gorm:"column:email_verified_at"`
	Password        string    `json:"-" gorm:"column:password;not null"`
	RememberToken   string    `json:"-" gorm:"column:remember_token;not null"`
	Role            int       `json:"role" gorm:"column:role"`
}

func (u UsersModel) TableName() string {
	return "users"
}
func (u *UsersModel) GetUserId() int {
	return int(u.ID)
}

func (u *UsersModel) GetEmail() string {
	return u.Email
}

func (u *UsersModel) GetRole() int {
	return u.Role
}

func (u *UsersModel) Mask(isAdmin bool) {
	u.GenUID(utils.DbTypeUser)
}

type UsersModelCreate struct {
	utils.SQLModel  `json:",inline"`
	Name            string     `json:"name" gorm:"not null;column:name"`
	Email           string     `json:"email" gorm:"column:email;unique"`
	Salt            string     `json:"-" gorm:"column:salt"`
	EmailVerifiedAt *time.Time `json:"-" gorm:"column:email_verified_at"`
	Password        string     `json:"password" gorm:"column:password;not null"`
	RememberToken   string     `json:"-" gorm:"column:remember_token;not null"`
	Role            int        `json:"-" gorm:"column:role"`
}

func (u *UsersModelCreate) TableName() string {
	return UsersModel{}.TableName()
}
func (u *UsersModelCreate) Mask(isAdmin bool) {
	u.GenUID(utils.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return UsersModel{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUserNameOrPassWordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
