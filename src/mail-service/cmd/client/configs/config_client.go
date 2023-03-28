package configs

import (
	"fmt"
	"os"
)

type Config struct {
	GrpcSendmailServiceTarget string
}

//var rootDir, _ = os.Getwd()

func LoadConfig() (config Config) {
	return Config{
		GrpcSendmailServiceTarget: fmt.Sprintf("localhost:%s", os.Getenv("SERVER_GRPC_PORT")),
	}
}
