package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDbInfo struct {
	connectionInfo
}

func newMysqlConnection(db INFO) IConnection {
	return &mysqlDbInfo{
		connectionInfo{
			Driver: "mysql",
			DbDns:  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.USER, db.PASSWORD, db.HOST, db.PORT, db.DBNAME),
		},
	}
}

func (mysqlConnect *mysqlDbInfo) ORM() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(mysqlConnect.DbDns), &gorm.Config{})
}
