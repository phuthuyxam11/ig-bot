package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/gin-validate-i18n-support/http_request"
	"igbot.com/authentication/component"
	"igbot.com/authentication/middleware"
	"igbot.com/authentication/modules/auth/request"
	"igbot.com/authentication/modules/auth/transport/ginauth"
)

func UserRoutesRegister(router *gin.Engine, appCtx component.AppContext) {
	routerAuth := router.Group("/auth")
	routerAuth.POST("/signup", ginauth.SignUp(appCtx))
	routerAuth.POST("/login", ginauth.SignIn(appCtx))
	routerAuth.GET("/profile", middleware.RequiredAuth(appCtx), ginauth.GetProfile(appCtx))
	routerAuth.GET("/verify-acc", http_request.GetParameterValidate[request.VerifyAccBody](appCtx.I18nBundle()), ginauth.VerifyProfile(appCtx))
	routerAuth.POST("/password-reset/verify", ginauth.PassWordRecoveryVerify(appCtx))
	routerAuth.GET("/password-reset", ginauth.PassWordRecovery(appCtx))
}
