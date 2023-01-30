package ginauth

import (
	"github.com/gin-gonic/gin"
	"igbot.com/authentication/component"
	"net/http"
)

func PassWordRecovery(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "sadasd")
	}
}

func PassWordRecoveryVerify(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "sadasd")
	}
}
