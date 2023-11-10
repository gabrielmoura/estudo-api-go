package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/glebarez/go-sqlite"
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
	Con = db
	return db
}
