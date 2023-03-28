package ginauth

import (
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component"
	"igbot.com/authentication/configs"
	userbiz "igbot.com/authentication/modules/auth/biz"
	userstorage "igbot.com/authentication/modules/auth/storage"
	"net/http"
)

func VerifyProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		config := configs.LoadConfig()
		queryParam := c.Request.URL.Query()
		store := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewAccountVerifyBusiness(store)
		expireTime := 60 * 60 * 24 * 2

		if config.EXPIREDTIMEVERIFYEMAIL > 0 {
			expireTime = config.EXPIREDTIMEVERIFYEMAIL
		}

		err := biz.Verify(c, appCtx, queryParam, expireTime)
		if err != nil {
			errResponse := common.NewErrorResponse(err, err.Error(), "", "validate_err")
			c.JSON(http.StatusBadRequest, errResponse)
		} else {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse("Account verified"))
		}
	}
}
