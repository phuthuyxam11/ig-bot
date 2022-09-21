package db

import (
	"database/sql"
	"gorm.io/gorm"
)

type IConnection interface {
	Connection() (*sql.DB, error)
	ORM() (*gorm.DB, error)
}

type INFO struct {
	USER     string
	PASSWORD string
	HOST     string
	PORT     string
	DBNAME   string
}
