package userbiz

import (
	"context"
	"errors"
	"igbot.com/authentication/component"
	usermod "igbot.com/authentication/modules/auth/model"
	"igbot.com/authentication/utils"
	"time"
)

type verifyAccStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermod.UserLoginData, error)
	UpdateUser(ctx context.Context, id int, data *usermod.UsersModelUpdate) error
}

type accountVerifyBusiness struct {
	verifyAccStorage verifyAccStorage
}

func NewAccountVerifyBusiness(storage verifyAccStorage) *accountVerifyBusiness {
	return &accountVerifyBusiness{
		verifyAccStorage: storage,
	}
}

func (business *accountVerifyBusiness) Verify(ctx context.Context, appCtx component.AppContext, params map[string][]string, expiredTime int) error {
	decodeUid, err := utils.FromBase58(params["id"][0])
	if err != nil {
		errMessage, err := utils.MessageRender(appCtx, "validate.verifyAcc.id.userexist", map[string]string{})
		if err != nil {
			return errors.New("user not found")
		}
		return errors.New(errMessage)
	}
	data, err := business.verifyAccStorage.FindUser(ctx, map[string]interface{}{"id": int(decodeUid.GetLocalID())})

	if !data.EmailVerifiedAt.IsZero() {
		errMessage, err := utils.MessageRender(appCtx, "validate.verifyAcc.activated", map[string]string{})
		if err != nil {
			return errors.New("user has been activated")
		}
		return errors.New(errMessage)
	}
	if data.EmailConfirmationToken != params["code"][0] {
		errMessage, err := utils.MessageRender(appCtx, "validate.verifyAcc.code.incorrect", map[string]string{})
		if err != nil {
			return errors.New("verify token incorrect")
		}
		return errors.New(errMessage)
	}
	if time.Now().Unix()-data.UpdatedAt.Unix() > int64(expiredTime) {
		errMessage, err := utils.MessageRender(appCtx, "validate.verifyAcc.code.expired", map[string]string{})
		if err != nil {
			return errors.New("the account registration code has expired ")
		}
		return errors.New(errMessage)
	}

	// update user verify acc
	userUpdate := usermod.UsersModelUpdate{EmailVerifiedAt: time.Now()}
	err = business.verifyAccStorage.UpdateUser(ctx, int(decodeUid.GetLocalID()), &userUpdate)
	return err
}
