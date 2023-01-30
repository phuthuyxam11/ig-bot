package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"igbot.com/authentication/component"
)

func I18nSupportMiddleWare(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Request.FormValue("lang")
		accept := c.Request.Header["Accept-Language"]
		localize := i18n.NewLocalizer(ac.I18nBundle(), language.English.String())
		if len(accept) > 0 {
			localize = i18n.NewLocalizer(ac.I18nBundle(), accept[0])
		}
		if len(lang) > 0 {
			localize = i18n.NewLocalizer(ac.I18nBundle(), lang)
		}
		ac.SetI18nLocalizer(localize)
		c.Next()
	}
}
