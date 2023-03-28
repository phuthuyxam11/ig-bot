package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type grpcConnectConfig struct {
	context context.Context
	target  string
	options []grpc.DialOption
}

func NewGrpcConnectConfig(ctx context.Context, target string, options []grpc.DialOption) *grpcConnectConfig {
	return &grpcConnectConfig{
		context: ctx,
		target:  target,
		options: options,
	}
}

func (cfg *grpcConnectConfig) GetClientConnect() (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(cfg.context, cfg.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error connect grpc server: ", err)
	}

	return conn, err
}
