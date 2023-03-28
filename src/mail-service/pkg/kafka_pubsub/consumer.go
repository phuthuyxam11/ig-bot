package kafka_pubsub

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type kafkaConfig struct {
	Topic            []string
	GroupId          string
	BootstrapServers string
	SecurityProtocol string
	//SaslMechanism    string `json:"sasl.mechanism"`
	//SaslUsername     string `json:"sasl.username"`
	//SaslPassword     string `json:"sasl.password"`
}

func NewKafkaConfig(topic []string, groupId string, boostrapServers string, securityProtocol string) *kafkaConfig {
	return &kafkaConfig{
		Topic:            topic,
		GroupId:          groupId,
		BootstrapServers: boostrapServers,
		SecurityProtocol: securityProtocol,
	}
}

func NewKafkaConsumer(config *kafkaConfig) (*kafka.Consumer, error) {
	fmt.Print("init kafka consumer, it may take a few seconds to init the connection\n")
	var kafkaConf = &kafka.ConfigMap{
		"api.version.request":       "true",
		"auto.offset.reset":         "latest",
		"heartbeat.interval.ms":     3000,
		"session.timeout.ms":        30000,
		"max.poll.interval.ms":      120000,
		"fetch.max.bytes":           1024000,
		"max.partition.fetch.bytes": 256000}
	if err := kafkaConf.SetKey("bootstrap.servers", config.BootstrapServers); err != nil {
		return nil, err
	}
	if err := kafkaConf.SetKey("group.id", config.GroupId); err != nil {
		return nil, err
	}
	switch config.SecurityProtocol {
	case "plaintext":
		if err := kafkaConf.SetKey("security.protocol", "plaintext"); err != nil {
			return nil, err
		}
	//case "sasl_ssl":
	//	kafkaConf.SetKey("security.protocol", "sasl_ssl")
	//	kafkaConf.SetKey("ssl.ca.location", "./conf/ca-cert.pem")
	//	kafkaConf.SetKey("sasl.username", cfg.SaslUsername)
	//	kafkaConf.SetKey("sasl.password", cfg.SaslPassword)
	//	kafkaConf.SetKey("sasl.mechanism", cfg.SaslMechanism)
	//case "sasl_plaintext":
	//	kafkaConf.SetKey("security.protocol", "sasl_plaintext")
	//	kafkaConf.SetKey("sasl.username", cfg.SaslUsername)
	//	kafkaConf.SetKey("sasl.password", cfg.SaslPassword)
	//	kafkaConf.SetKey("sasl.mechanism", cfg.SaslMechanism)

	default:
		panic(kafka.NewError(kafka.ErrUnknownProtocol, "unknown protocol", true))
	}

	consumer, err := kafka.NewConsumer(kafkaConf)

	//defer func(consumer *kafka.Consumer) {
	//	err := consumer.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(consumer)

	if err != nil {
		return nil, err
	}

	fmt.Print("init kafka consumer success\n")
	return consumer, nil
}
