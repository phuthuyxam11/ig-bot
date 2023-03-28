package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"
	"igbot.com/sendmail/api/pb"
	"igbot.com/sendmail/cmd/client/configs"
	"igbot.com/sendmail/pkg/decode"
	"igbot.com/sendmail/pkg/helpers"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	rootDir, _ := os.Getwd()
	pemServerCA, err := ioutil.ReadFile(rootDir + "/ssl/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("fail to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}
func SendMailService(data *decode.VerifyMailTopic[decode.VerifyMailTopicData]) error {
	cfg := configs.LoadConfig()
	credential, err := LoadTLSCredentials()
	if err != nil {
		log.Fatalln("Error create client credential: ", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := []grpc.DialOption{grpc.WithTransportCredentials(credential)}
	grpcConnectConfig := NewGrpcConnectConfig(ctx, cfg.GrpcSendmailServiceTarget, options)

	connect, err := grpcConnectConfig.GetClientConnect()
	defer connect.Close()

	if err != nil {
		log.Fatalln("Fail to connect grpc server: ", err)
		return err
	}
	client := pb.NewMailSenderServiceClient(connect)

	_, err = client.SendMailVerifyAcc(ctx, &pb.MailVerifyAccRequest{
		ActiveUrl: data.Value.VerifyUrl,
		EmailInfo: &pb.EmailInfo{
			Subject: "Local send mail verify",
			To:      helpers.StringSliceToMap(data.Value.To),
			From:    "noreply@igbot.com",
			Cc:      nil,
			Bcc:     nil,
			Html:    true,
			Lang:    data.Value.Lang,
		},
		CreateAt: timestamppb.New(time.Now()),
	})

	return err
}
