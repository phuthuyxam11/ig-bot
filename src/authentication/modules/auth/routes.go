package auth

import (
	"github.com/gin-gonic/gin"
	"igbot.com/authentication/component"
	"igbot.com/authentication/modules/auth/transport/ginauth"
)

func UserRoutesRegister(router *gin.Engine, appCtx component.AppContext) {
	routerAuth := router.Group("/auth")
	//routerAuth.POST("/signup", ginauth.SignUp(appCtx))
	routerAuth.GET("/login", ginauth.SignIn(appCtx))
	//routerAuth.GET("/profile", middleware.RequiredAuth(appCtx), ginauth.GetProfile(appCtx))
}
