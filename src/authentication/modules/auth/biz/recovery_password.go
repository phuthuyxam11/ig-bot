package userbiz

import (
	"context"
	"errors"
	"igbot.com/authentication/component"
	"igbot.com/authentication/configs"
	"igbot.com/authentication/kafka_pubsub"
	usermod "igbot.com/authentication/modules/auth/model"
	"igbot.com/authentication/utils"
)

type recoveryPassWordStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermod.UserLoginData, error)
}

type recoveryPassWordBusiness struct {
	recoveryPassWordStorage recoveryPassWordStorage
	msgBroker               kafka_pubsub.RegisterUserMessageBrokerRepository
}
type RecoveryPasswordMessage struct {
	RecoveryUrl string
}

func NewRecoveryPassWord(storage recoveryPassWordStorage,
	msgBroker kafka_pubsub.RegisterUserMessageBrokerRepository) *recoveryPassWordBusiness {
	return &recoveryPassWordBusiness{
		recoveryPassWordStorage: storage,
		msgBroker:               msgBroker,
	}
}

func (business *recoveryPassWordBusiness) SendRecoveryMail(ctx context.Context, appCtx component.AppContext, email string) error {
	data, err := business.recoveryPassWordStorage.FindUser(ctx, map[string]interface{}{"email": email})

	if data == nil {
		return errors.New("Recovery password: email " + email + " not found")
	}
	if err != nil {
		return err
	}

	if !data.EmailVerifiedAt.IsZero() {
		errMessage, err := utils.MessageRender(appCtx, "validate.verifyAcc.activated", map[string]string{})
		if err != nil {
			return errors.New("user has been activated")
		}
		return errors.New(errMessage)
	}
	config := configs.LoadConfig()

	if err := business.msgBroker.RecoveryPasswordEmailMessage(ctx, RecoveryPasswordMessage{RecoveryUrl: "#"}); err != nil {
		return kafka_pubsub.NewErrorKafkaSendMessage(err,
			config.KAFKA_TOPIC_SEND_MAIL,
			"acc.recovery.sendmail")
	}

	return err
}
