package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"igbot.com/sendmail/api/pb"
	"igbot.com/sendmail/cmd/client"
	"igbot.com/sendmail/cmd/server/configs"
	"igbot.com/sendmail/component"
	"igbot.com/sendmail/internal/middleware"
	"igbot.com/sendmail/internal/services"
	"igbot.com/sendmail/pkg/db"
	"igbot.com/sendmail/pkg/decode"
	"igbot.com/sendmail/pkg/kafka_pubsub"
	"log"
	"net"
)

func runGrpcServer(appContext component.AppContext) {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(middleware.RecoverInterceptor))),
	)
	serverInstance := services.NewMailSenderServer(appContext)

	pb.RegisterMailSenderServiceServer(grpcServer, serverInstance)
	reflection.Register(grpcServer)
	conf := appContext.GetAppConfig()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCSERVERPORT))
	if err != nil {
		log.Fatal("cannot create listener")
	}
	log.Println("===== grpc start with port: ", conf.GRPCSERVERPORT, " =====")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func KafkaConsumerServe() {
	config := configs.LoadConfig()
	cfg := kafka_pubsub.NewKafkaConfig([]string{config.KAFKA_TOPIC_SEND_MAIL}, config.KAFKA_CONSUMER_GROUP_SEND_MAIL, config.SPRING_KAFKA_BOOTSTRAP_SERVERS, "plaintext")
	consumer, err := kafka_pubsub.NewKafkaConsumer(cfg)
	if err != nil {
		return
	}

	if err := consumer.SubscribeTopics(cfg.Topic, nil); err != nil {
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			var res decode.VerifyMailTopic[decode.VerifyMailTopicData]

			data, err := decode.SendMailTopicConverter[decode.VerifyMailTopic[decode.VerifyMailTopicData]](msg.Value, &res)

			if err != nil {
				fmt.Println(err)
				return
			}

			switch data.Type {
			case config.KAFKA_EVENT_TYPE["email_register"]:
				// send mail register logic
				fmt.Println("send mail register")
				err := client.SendMailService(data)
				if err != nil {
					log.Printf("Send mail error: %v", err)
					return
				}
			case config.KAFKA_EVENT_TYPE["email_recovery"]:
				// send mail recovery logic
				//err := client.SendMailService()
				//if err != nil {
				//	return
				//}

			default:
				log.Printf("Kafka event type: %s not support", data.Type)
			}

			fmt.Println(data.Type, data.Value)
			_, err = consumer.CommitMessage(msg)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			log.Printf("consumer error: %v (%v)\n", err, msg)
			continue
		}
	}
	consumer.Close()
}

func main() {
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

	appCtx := component.NewAppContext(orm, config)

	go runGrpcServer(appCtx)

	kafkaServe := make(chan interface{})
	go func() {
		KafkaConsumerServe()
		kafkaServe <- 1
	}()
	<-kafkaServe

}
