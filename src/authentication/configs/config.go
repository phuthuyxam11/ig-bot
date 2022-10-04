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
}

//DBDrive    string `mapstructure:"DB_DRIVER"`
//DBUSER     string `mapstructure:"DB_USER"`
//DBPASSWORD string `mapstructure:"DB_PASSWORD"`
//DBHOST     string `mapstructure:"DB_HOST"`
//DBPORT     int    `mapstructure:"DB_PORT"`
//DBNAME     string `mapstructure:"DB_NAME"`
//SERVERPOST string `mapstructure:"SERVER_POST"`

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
	}
}
