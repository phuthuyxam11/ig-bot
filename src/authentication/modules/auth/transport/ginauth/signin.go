package ginauth

import (
	"github.com/gin-gonic/gin"
	"igbot.com/authentication/component"
	"igbot.com/authentication/component/hasher"
	"igbot.com/authentication/component/tokenprovider/jwt"
	userbiz "igbot.com/authentication/modules/auth/biz"
	usermod "igbot.com/authentication/modules/auth/model"
	userstorage "igbot.com/authentication/modules/auth/storage"
	"net/http"
)

type uj struct {
	Na string
}

func SignIn(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermod.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		md5 := hasher.NewMd5Hash()
		signInStorage := userstorage.NewSqlStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		biz := userbiz.NewLoginBusiness(signInStorage, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, account)
	}
}
