package ginauth

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/gin-validate-i18n-support/http_request"
	"igbot.com/authentication/component"
	"igbot.com/authentication/configs"
	"igbot.com/authentication/kafka_pubsub"
	userbiz "igbot.com/authentication/modules/auth/biz"
	"igbot.com/authentication/modules/auth/request"
	userstorage "igbot.com/authentication/modules/auth/storage"
	"net/http"
)

func PassWordRecovery(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data request.RecoveryPasswordVerifyBody
		if err := c.ShouldBind(&data); err != nil {
			http_request.ValidationRender[request.RecoveryPasswordVerifyBody](c, err, data, appCtx.I18nBundle())
			return
		}

		store := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
		config := configs.LoadConfig()
		producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.SPRING_KAFKA_BOOTSTRAP_SERVERS})
		if err != nil {
			fmt.Printf("erroror with create producer")
			panic(err)
		}
		publisher := kafka_pubsub.NewPublisher(producer, config.KAFKA_TOPIC_SEND_MAIL, "acc.recovery.sendmail")

		biz := userbiz.NewRecoveryPassWord(store, publisher)
		err = biz.SendRecoveryMail(c, appCtx, data.Email)

		if err != nil {
			// send mail
			panic(err)
		}

		c.JSON(http.StatusOK, "ok!")
	}
}
