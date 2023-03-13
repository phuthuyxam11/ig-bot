package userbiz

import (
	"context"
	"errors"
	"fmt"
	"igbot.com/authentication/component"
	usermod "igbot.com/authentication/modules/auth/model"
	"igbot.com/authentication/utils"
)

type recoveryPassWordStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermod.UserLoginData, error)
}

type recoveryPassWordBusiness struct {
	recoveryPassWordStorage recoveryPassWordStorage
}

func NewRecoveryPassWord(storage recoveryPassWordStorage) *recoveryPassWordBusiness {
	return &recoveryPassWordBusiness{
		recoveryPassWordStorage: storage,
	}
}

func (business *recoveryPassWordBusiness) SendRecoveryMail(ctx context.Context, appCtx component.AppContext, email string) error {
	data, err := business.recoveryPassWordStorage.FindUser(ctx, map[string]interface{}{"email": email})

	if data == nil {
		return errors.New("Recovery password: email " + email + " not found")
	}
	if err != nil {
		return err
	}

	if !data.EmailVerifiedAt.IsZero() {
		errMessage, err := utils.MessageRender(appCtx, "validate.verifyAcc.activated", map[string]string{})
		if err != nil {
			return errors.New("user has been activated")
		}
		return errors.New(errMessage)
	}
	fmt.Println("ASDASDD")
	return err
}
