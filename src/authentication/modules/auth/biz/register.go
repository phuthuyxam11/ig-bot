package userbiz

import (
	"200lab/struct/common"
	userenum "200lab/struct/modules/auth/enum"
	usermodel "200lab/struct/modules/auth/model"

	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.UsersModel, error)
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

	salt := common.GenSalt(50)

	data.Password = business.hashed.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = userenum.GUEST

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
