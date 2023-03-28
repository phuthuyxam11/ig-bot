package userbiz

import (
	"context"
	"fmt"
	"github.com/phuthuyxam11/go-common-service/common"
	"igbot.com/authentication/configs"
	"igbot.com/authentication/kafka_pubsub"
	usermodel "igbot.com/authentication/modules/auth/model"
	"igbot.com/authentication/utils"
	"strconv"
	"time"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.UserLoginData, error)
	CreateUser(ctx context.Context, data *usermodel.UsersModelCreate) error
}

type Hashed interface {
	Hash(data string) string
}

type sendMailRegister struct {
	VerifyUrl string   `json:"verifyUrl"`
	To        []string `json:"to"`
	Lang      string   `json:"lang"`
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hashed          Hashed
	msgBroker       kafka_pubsub.RegisterUserMessageBrokerRepository
}

func NewRegisterBusiness(registerStorage RegisterStorage, hashed Hashed, msgBroker kafka_pubsub.RegisterUserMessageBrokerRepository) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hashed:          hashed,
		msgBroker:       msgBroker,
	}
}

func (business *registerBusiness) Register(ctx context.Context, lang string, data *usermodel.UsersModelCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	config := configs.LoadConfig()

	salt := utils.GenSalt(50)
	emailConfirmationToken := utils.GenSalt(200) + strconv.FormatInt(time.Now().Unix(), 10)
	tokenGenerationTime := config.TIMELIMITREGISTERTOKEN

	data.Password = business.hashed.Hash(data.Password + salt)
	data.PasswordSalt = salt
	data.EmailConfirmationToken = emailConfirmationToken
	data.TokenGenerationTime = tokenGenerationTime

	//data.Role = userenum.GUEST

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	data.Mask(false)
	verifyUrl := fmt.Sprintf("%s?code=%s&id=%s", utils.GetRoutersMap()["auth_verify_acc"], emailConfirmationToken, data.FakeId.String())

	if err := business.msgBroker.VerifyEmailMessage(ctx,
		sendMailRegister{
			VerifyUrl: verifyUrl,
			To:        []string{data.Email},
			Lang:      lang},
	); err != nil {
		return kafka_pubsub.NewErrorKafkaSendMessage(err, config.KAFKA_TOPIC_SEND_MAIL, "acc.register.sendmail")
	}

	return nil
}
