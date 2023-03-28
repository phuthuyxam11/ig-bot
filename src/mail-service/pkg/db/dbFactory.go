package db

import (
	"fmt"
	"igbot.com/sendmail/cmd/server/configs"
)

func GetConnection() (IConnection, error) {
	config := configs.LoadConfig()

	if config.DBDrive == "mysql" {
		return newMysqlConnection(INFO{
			USER:     config.DBUSER,
			PASSWORD: config.DBPASSWORD,
			HOST:     config.DBHOST,
			PORT:     config.DBPORT,
			DBNAME:   config.DBNAME,
		}), nil
	}
	if config.DBDrive == "postgres" {
		return newPostgresConnection(INFO{
			USER:     config.DBUSER,
			PASSWORD: config.DBPASSWORD,
			HOST:     config.DBHOST,
			PORT:     config.DBPORT,
			DBNAME:   config.DBNAME,
		}), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
