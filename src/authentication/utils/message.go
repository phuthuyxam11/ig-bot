package utils

import (
	"errors"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"igbot.com/authentication/component"
)

func MessageRender(appContext component.AppContext, key string, templateData map[string]string) (string, error) {
	if appContext.GetI18nLocalizer() == nil {
		return "", errors.New("app context missing I18nLocalizer config")
	}
	localizeValidateMessage := i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: templateData,
	}
	message, err := appContext.GetI18nLocalizer().Localize(&localizeValidateMessage)
	fmt.Println(appContext.GetI18nLocalizer())
	if err != nil {
		return "", errors.New("message render fail")
	}
	return message, nil
}
