package ginauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"igbot.com/authentication/component"
	"igbot.com/authentication/modules/auth/request"
	"net/http"
)

func PassWordRecovery(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data request.RecoveryPasswordVerifyBody
		if err := c.ShouldBind(&data); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
			//return
			for _, fieldErr := range err.(validator.ValidationErrors) {
				fmt.Println(fieldErr)
			}
		}
		//if err := c.ShouldBindBodyWith(&data, binding.J); err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(data) // return nil
		//store := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
		//biz := userbiz.NewRecoveryPassWord(store)
		//err := biz.SendRecoveryMail(c, appCtx, data.Email)
		//
		//if err != nil {
		//	// send mail
		//	fmt.Println("send mail service...")
		//}

		c.JSON(http.StatusOK, "Send mail")
	}
}
