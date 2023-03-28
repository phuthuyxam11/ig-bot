package component

import (
	"gorm.io/gorm"
	"igbot.com/sendmail/cmd/server/configs"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetAppConfig() configs.Config
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	configs   configs.Config
}

func NewAppContext(db *gorm.DB, configs configs.Config) *appCtx {
	return &appCtx{db: db, configs: configs}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetAppConfig() configs.Config {
	return ctx.configs
}
func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
