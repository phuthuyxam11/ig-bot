package auth

import (
	"context"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"igbot.com/authentication/component"
	userbiz "igbot.com/authentication/modules/auth/biz"
	userstorage "igbot.com/authentication/modules/auth/storage"
	"igbot.com/authentication/utils"
)

func RegisterValidation(appCtx component.AppContext) error {
	ctx := context.Background()
	// getting the validation engine and type casting it.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// registering validation for userexist
		err := v.RegisterValidation("userexist", func(fl validator.FieldLevel) bool {
			// convert field value to string
			value := fl.Field().String()
			decodeUid, err := utils.FromBase58(value)
			if err != nil {
				return false
			}
			store := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
			biz := userbiz.NewGetUserBiz(store)
			user, err := biz.GetUserBiz(ctx, int(decodeUid.GetLocalID()))
			if err != nil {
				return false
			}

			if user.GetUserId() > 0 {
				return true
			}
			return false
		})

		if err != nil {
			return err
		}
	}
	return nil
}
