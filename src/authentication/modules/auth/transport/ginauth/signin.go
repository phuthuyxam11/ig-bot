package ginauth

import (
	"github.com/gin-gonic/gin"
	"igbot.com/authentication/component"
	"net/http"
)

type uj struct {
	Na string
}

func SignIn(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDBConnection()
		//var data usermod.UserLogin
		//if err := c.ShouldBind(&data); err != nil {
		//	panic(err)
		//}
		//md5 := hasher.NewMd5Hash()
		//signInStorage := userstorage.NewSqlStore(db)
		//tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		//biz := userbiz.NewLoginBusiness(signInStorage, tokenProvider, md5, 60*60*24*30)
		//account, err := biz.Login(c.Request.Context(), &data)
		//
		//if err != nil {
		//	panic(err)
		//}

		c.JSON(http.StatusOK, uj{Na: "sdddd"})
	}
}
