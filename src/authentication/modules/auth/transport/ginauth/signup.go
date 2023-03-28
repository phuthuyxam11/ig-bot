package ginauth

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/phuthuyxam11/go-common-service/common"
	"golang.org/x/text/language"
	"igbot.com/authentication/component"
	"igbot.com/authentication/component/hasher"
	"igbot.com/authentication/configs"
	"igbot.com/authentication/kafka_pubsub"
	userbiz "igbot.com/authentication/modules/auth/biz"
	usermod "igbot.com/authentication/modules/auth/model"
	userstorage "igbot.com/authentication/modules/auth/storage"
	"net/http"
)

func SignUp(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermod.UsersModelCreate
		lang, exist := c.Get("lang_accept")
		if !exist {
			lang = language.English.String()
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		config := configs.LoadConfig()

		producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.SPRING_KAFKA_BOOTSTRAP_SERVERS})
		if err != nil {
			fmt.Printf("erroror with create producer")
			panic(err)
		}
		publisher := kafka_pubsub.NewPublisher(producer, config.KAFKA_TOPIC_SEND_MAIL, "acc.register.sendmail")

		biz := userbiz.NewRegisterBusiness(store, md5, publisher)
		if err := biz.Register(c.Request.Context(), fmt.Sprintf("%v", lang), &data); err != nil {
			fmt.Println(err)
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
