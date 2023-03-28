package kafka_pubsub

import (
	"context"
	"fmt"
	"github.com/phuthuyxam11/go-common-service/common"
	"log"
	"net/http"
)

type RegisterUserMessageBrokerRepository interface {
	VerifyEmailMessage(ctx context.Context, data interface{}) error
	RecoveryPasswordEmailMessage(ctx context.Context, data interface{}) error
}

func NewErrorKafkaSendMessage(root error, topic string, topicKey string) *common.AppError {
	log.Println("Error when send message: ", root)
	return &common.AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    fmt.Sprintf("can't send message to kafka topic: %s, topicKey: %s", topic, topicKey),
		Log:        fmt.Sprintf("can't send message to kafka topic: %s, topicKey: %s", topic, topicKey),
		Key:        "err_send_kafka",
	}
}
