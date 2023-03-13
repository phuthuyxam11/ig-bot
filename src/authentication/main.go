package main

import (
	"fmt"
	"igbot.com/authentication/modules"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/gin-validate-i18n-support/utils"
	"igbot.com/authentication/component"
	"igbot.com/authentication/configs"
	"igbot.com/authentication/db"
	"igbot.com/authentication/middleware"
	"igbot.com/authentication/modules/auth"
)

// @title Todo Application
// @description This is a todo list management application
// @version 1.0
// @host localhost:8081
// @BasePath /api/v1

func main() {

	// load config
	config := configs.LoadConfig()

	connection, err := db.GetConnection()
	orm, err := connection.ORM()

	if err != nil {
		log.Fatalln(err)
	}

	// migration db
	err = db.Migrate(orm)
	if err != nil {
		log.Fatalln(err)
	}

	// register i18nSupport
	i18nBundle, err := utils.I18nInit(config.LANGUAGEFOLDERPATH)
	if err != nil {
		log.Fatalln(err)
	}

	appCtx := component.NewAppContext(orm, i18nBundle)

	// register custom validation
	err = auth.RegisterValidation(appCtx)
	if err != nil {
		log.Fatalln(err)
	}

	// routes
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.I18nSupportMiddleWare(appCtx))

	// authentication routes
	modules.UserRoutesRegister(r, appCtx)

	errServer := http.ListenAndServe(":"+config.SERVERPOST, r)
	if errServer != nil {
		fmt.Println(errServer)
		return
	}
}
