package configs

import (
	"os"
	"time"
)

type Config struct {
	DBDrive                        string
	DBUSER                         string
	DBPASSWORD                     string
	DBHOST                         string
	DBPORT                         string
	DBNAME                         string
	SERVERPOST                     string
	TIMELIMITREGISTERTOKEN         time.Time
	EXPIREDTIMEVERIFYEMAIL         int
	TEMPLATEFOLDERPATH             string
	GRPCSERVERPORT                 string
	KAFKA_TOPIC_SEND_MAIL          string
	SPRING_KAFKA_BOOTSTRAP_SERVERS string
	KAFKA_CONSUMER_GROUP_SEND_MAIL string
	KAFKA_EVENT_TYPE               map[string]string
}

var rootDir, _ = os.Getwd()

func LoadConfig() (config Config) {
	eventType := map[string]string{
		"email_register": "account.register.sendMail",
		"email_recovery": "acc.recovery.sendmail",
	}
	return Config{
		DBDrive:                        os.Getenv("DB_DRIVER"),
		DBUSER:                         os.Getenv("DB_USER"),
		DBPASSWORD:                     os.Getenv("DB_PASSWORD"),
		DBHOST:                         os.Getenv("DB_HOST"),
		DBPORT:                         os.Getenv("DB_PORT"),
		DBNAME:                         os.Getenv("DB_NAME"),
		SERVERPOST:                     os.Getenv("SERVER_POST"),
		TIMELIMITREGISTERTOKEN:         time.Now().Add(time.Second * 60 * 60 * 24 * 3),
		EXPIREDTIMEVERIFYEMAIL:         60 * 60 * 24 * 4,
		TEMPLATEFOLDERPATH:             rootDir + "/pkg/mailer/template/",
		GRPCSERVERPORT:                 os.Getenv("SERVER_GRPC_PORT"),
		KAFKA_TOPIC_SEND_MAIL:          os.Getenv("KAFKA_TOPIC_SEND_MAIL"),
		SPRING_KAFKA_BOOTSTRAP_SERVERS: os.Getenv("SPRING_KAFKA_BOOTSTRAP_SERVERS"),
		KAFKA_CONSUMER_GROUP_SEND_MAIL: "send_mail_group",
		KAFKA_EVENT_TYPE:               eventType,
	}
}
