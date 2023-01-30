package configs

import (
	"os"
	"time"
)

type Config struct {
	DBDrive                string
	DBUSER                 string
	DBPASSWORD             string
	DBHOST                 string
	DBPORT                 string
	DBNAME                 string
	SERVERPOST             string
	TIMELIMITREGISTERTOKEN time.Time
	EXPIREDTIMEVERIFYEMAIL int
	LANGUAGEFOLDERPATH     string
}

var rootDir, _ = os.Getwd()

func LoadConfig() (config Config) {
	return Config{
		DBDrive:                os.Getenv("DB_DRIVER"),
		DBUSER:                 os.Getenv("DB_USER"),
		DBPASSWORD:             os.Getenv("DB_PASSWORD"),
		DBHOST:                 os.Getenv("DB_HOST"),
		DBPORT:                 os.Getenv("DB_PORT"),
		DBNAME:                 os.Getenv("DB_NAME"),
		SERVERPOST:             os.Getenv("SERVER_POST"),
		TIMELIMITREGISTERTOKEN: time.Now().Add(time.Second * 60 * 60 * 24 * 3),
		EXPIREDTIMEVERIFYEMAIL: 60 * 60 * 24 * 4,
		LANGUAGEFOLDERPATH:     rootDir + "/lang",
	}
}
