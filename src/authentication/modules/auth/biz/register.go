package userbiz

import (
	"context"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/configs"
	usermodel "igbot.com/authentication/modules/auth/model"
	"igbot.com/authentication/utils"
	"time"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.UserLoginData, error)
	CreateUser(ctx context.Context, data *usermodel.UsersModelCreate) error
}

type Hashed interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hashed          Hashed
}

func NewRegisterBusiness(registerStorage RegisterStorage, hashed Hashed) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hashed:          hashed,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UsersModelCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	config := configs.LoadConfig()

	salt := utils.GenSalt(50)
	emailConfirmationToken := utils.GenSalt(200) + time.Now().String()
	tokenGenerationTime := config.TIMELIMITREGISTERTOKEN

	data.Password = business.hashed.Hash(data.Password + salt)
	data.PasswordSalt = salt
	data.EmailConfirmationToken = emailConfirmationToken
	data.TokenGenerationTime = tokenGenerationTime

	//data.Role = userenum.GUEST

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
