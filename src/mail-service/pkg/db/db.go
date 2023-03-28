package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type connectionInfo struct {
	Driver string
	DbDns  string
}

func (db *connectionInfo) Connection() (*sql.DB, error) {
	connect, err := sql.Open(db.Driver, db.DbDns)
	if err != nil {
		log.Fatal("Connect db fail: ", err)
		return nil, err
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal("Connect db fail: ", err)
		}
	}(connect)

	err = connect.Ping()
	if err != nil {
		log.Fatal("Ping db fail: ", err)
		return nil, err
	}

	return connect, nil
}
