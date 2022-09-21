package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"igbot.com/authentication/component"
	"igbot.com/authentication/configs"
	"igbot.com/authentication/db"
	"igbot.com/authentication/middleware"
	"igbot.com/authentication/modules/auth"
	"log"
	"net/http"
)

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

	appCtx := component.NewAppContext(orm)
	// routes
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	// authentication routes
	auth.UserRoutesRegister(r, appCtx)

	errServer := http.ListenAndServe(":"+config.SERVERPOST, r)
	if errServer != nil {
		fmt.Println(errServer)
		return
	}
}
