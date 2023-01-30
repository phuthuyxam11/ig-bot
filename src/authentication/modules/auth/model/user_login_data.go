package usermodel

import (
	"errors"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component/tokenprovider"
	"igbot.com/authentication/utils"
	"time"
)

const EntityName = "User"

type UserLoginData struct {
	utils.SQLModel         `json:",inline"`
	LoginName              string    `json:"name" gorm:"not null;column:login_name"`
	Password               string    `json:"-" gorm:"column:password;not null"`
	PasswordSalt           string    `json:"-" gorm:"column:password_salt"`
	Email                  string    `json:"email" gorm:"column:email;unique"`
	EmailConfirmationToken string    `json:"-" gorm:"column:email_confirmation_token"`
	TokenGenerationTime    time.Time `json:"-" gorm:"column:token_generation_time"`
	EmailVerifiedAt        time.Time `json:"-" gorm:"column:email_verified_at"`
	RememberToken          string    `json:"-" gorm:"column:remember_token"`
	PasswordRecoveryToken  string    `json:"-" gorm:"column:password_recovery_token"`
	RecoveryTokenTime      time.Time `json:"-" gorm:"column:recovery_token_time"`
}

func (u UserLoginData) TableName() string {
	return "user_login_data"
}

func (u *UserLoginData) GetUserId() int {
	return int(u.ID)
}

func (u *UserLoginData) GetEmail() string {
	return u.Email
}

//func (u *UserLoginData) GetRole() int {
//	return u.
//}

func (u *UserLoginData) Mask(isAdmin bool) {
	u.GenUID(utils.DbTypeUser)
}

type UsersModelCreate struct {
	utils.SQLModel         `json:",inline"`
	LoginName              string    `json:"name" gorm:"not null;column:login_name"`
	Email                  string    `json:"email" gorm:"column:email;unique"`
	Password               string    `json:"password" gorm:"column:password;not null"`
	PasswordSalt           string    `json:"-" gorm:"column:password_salt"`
	RememberToken          string    `json:"-" gorm:"column:remember_token;not null"`
	EmailConfirmationToken string    `json:"-" gorm:"column:email_confirmation_token;not null"`
	TokenGenerationTime    time.Time `json:"-" gorm:"column:token_generation_time;not null"`
}

func (u *UsersModelCreate) TableName() string {
	return UserLoginData{}.TableName()
}
func (u *UsersModelCreate) Mask(isAdmin bool) {
	u.GenUID(utils.DbTypeUser)
}

type UsersModelUpdate struct {
	utils.SQLModel         `json:",inline"`
	LoginName              string    `json:"name" gorm:"not null;column:login_name;unique"`
	Email                  string    `json:"email" gorm:"column:email;unique"`
	Password               string    `json:"password" gorm:"column:password;not null"`
	PasswordSalt           string    `json:"-" gorm:"column:password_salt"`
	RememberToken          string    `json:"-" gorm:"column:remember_token;not null"`
	EmailConfirmationToken string    `json:"-" gorm:"column:email_confirmation_token;not null"`
	TokenGenerationTime    time.Time `json:"-" gorm:"column:token_generation_time;not null"`
	EmailVerifiedAt        time.Time `json:"-" gorm:"column:email_verified_at"`
}

func (u *UsersModelUpdate) TableName() string {
	return UserLoginData{}.TableName()
}
func (u *UsersModelUpdate) Mask(isAdmin bool) {
	u.GenUID(utils.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return UserLoginData{}.TableName()
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

	ErrAccIsNotVerify = common.NewCustomError(
		errors.New("account is not verified email. plz check your email"),
		"account is not verified email. plz check your email",
		"ErrAccIsNotVerify",
	)
)
