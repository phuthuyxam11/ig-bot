package ginauth

import (
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component"
	"net/http"
)

func SignUp(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDBConnection()
		//var data usermod.UsersModelCreate
		//
		//if err := c.ShouldBind(&data); err != nil {
		//	panic(err)
		//}
		//store := userstorage.NewSqlStore(db)
		//md5 := hasher.NewMd5Hash()
		//biz := userbiz.NewRegisterBusiness(store, md5)
		//if err := biz.Register(c.Request.Context(), &data); err != nil {
		//	panic(err)
		//}
		//data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("asdasdasd"))
	}
}