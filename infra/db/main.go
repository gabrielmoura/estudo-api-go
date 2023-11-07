package db

import (
	"database/sql"
	"errors"
	_ "github.com/glebarez/go-sqlite"
	"log"
)

var (
	Con         *sql.DB
	ErrToInsert = errors.New("Erro ao inserir")
)

func Conn(driverName, dataSource string) *sql.DB {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
