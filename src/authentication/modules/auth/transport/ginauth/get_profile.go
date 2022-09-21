package ginauth

import (
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component"
	"net/http"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("asdasd"))
	}
}
