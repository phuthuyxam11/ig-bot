package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgDbInfo struct {
	connectionInfo
}

func newPostgresConnection(db INFO) IConnection {
	return &pgDbInfo{
		connectionInfo{
			Driver: "postgres",
			DbDns:  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.HOST, db.PORT, db.USER, db.PASSWORD, db.DBNAME),
		},
	}
}

func (pgc *pgDbInfo) ORM() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(pgc.DbDns), &gorm.Config{})
}
