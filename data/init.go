package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func init() {
	connect()
}

func connect() {
	dsn := "postgres://postgres:postgres@localhost/notedb?sslmode=disable&binary_parameters=yes"

	var err error
	Db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}
