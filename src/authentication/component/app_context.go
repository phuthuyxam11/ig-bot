package component

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	// UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	// GetPubsub() pubsub.Pubsub
	I18nBundle() *i18n.Bundle
	SetI18nLocalizer(localize *i18n.Localizer)
	GetI18nLocalizer() *i18n.Localizer
}

type appCtx struct {
	db *gorm.DB
	// upProvider uploadprovider.UploadProvider
	secretKey string
	// pb         pubsub.Pubsub
	i18nBundle    *i18n.Bundle
	i18nLocalizer *i18n.Localizer
}

func NewAppContext(db *gorm.DB, bundle *i18n.Bundle) *appCtx {
	return &appCtx{db: db, i18nBundle: bundle, i18nLocalizer: nil}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) I18nBundle() *i18n.Bundle {
	return ctx.i18nBundle
}

func (ctx *appCtx) SetI18nLocalizer(localize *i18n.Localizer) {
	ctx.i18nLocalizer = localize
}

func (ctx *appCtx) GetI18nLocalizer() *i18n.Localizer {
	return ctx.i18nLocalizer
}

// func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
// 	return ctx.upProvider
// }

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

// func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
// 	return ctx.pb
// }
