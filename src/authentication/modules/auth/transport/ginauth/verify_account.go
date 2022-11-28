package ginauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/component"
	"net/http"
)

func VerifyProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate
		//validate := middleware.NewRequest(c, VerifyAccBody{})
		//utils.Validate(validate)
		//println(validateErr)
		//if validateErr != nil {
		//	c.AbortWithError(http.StatusBadRequest, validateErr)
		//}
		fmt.Println(c.Request.URL.Query())
		// check verify
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ád"))
	}
}
