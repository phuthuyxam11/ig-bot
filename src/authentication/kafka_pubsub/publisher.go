package kafka_pubsub

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type event struct {
	Type  string
	Value interface{}
}

type Publisher struct {
	producer  *kafka.Producer
	topicName string
	key       string
}

func NewPublisher(producer *kafka.Producer, topicName string, key string) *Publisher {
	return &Publisher{
		producer:  producer,
		topicName: topicName,
		key:       key,
	}
}

func (p *Publisher) publish(ctx context.Context, msgType string, data interface{}) error {
	var b bytes.Buffer

	evt := event{
		Type:  msgType,
		Value: data,
	}

	if err := json.NewEncoder(&b).Encode(evt); err != nil {
		return err
	}

	if err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topicName,
			Partition: kafka.PartitionAny,
		},
		Value: b.Bytes(),
		Key:   []byte(p.key),
	}, nil); err != nil {
		return err
	}

	return nil
}

func (p *Publisher) VerifyEmailMessage(ctx context.Context, data interface{}) error {
	return p.publish(ctx, "account.register.sendMail", data)
}
func (p *Publisher) RecoveryPasswordEmailMessage(ctx context.Context, data interface{}) error {
	return p.publish(ctx, "account.recoveryPassword.sendMail", data)
}
