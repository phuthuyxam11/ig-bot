package ginauth

import (
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component"
	"net/http"
)

func VerifyProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		// check verify
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ád"))
	}
}
