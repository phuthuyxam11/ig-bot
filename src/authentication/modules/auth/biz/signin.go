package userbiz

import (
	"context"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component"
	"igbot.com/authentication/component/tokenprovider"
	usermod "igbot.com/authentication/modules/auth/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermod.UserLoginData, error)
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hashed
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hashed, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (business *loginBusiness) Login(ctx context.Context, data *usermod.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermod.ErrUserNameOrPassWordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.PasswordSalt)

	if user.Password != passHashed {
		return nil, usermod.ErrUserNameOrPassWordInvalid
	}

	if user.EmailVerifiedAt.IsZero() {
		return nil, usermod.ErrAccIsNotVerify
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.GetUserId(),
		Role:   "guest",
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := business.tokenProvider.Generate(payload, business.tkCfg.GetRtExp())
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	//account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
}
