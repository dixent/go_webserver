package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Connection *sqlx.DB

func InitConnection() (*sqlx.DB) {
	conn, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=go_webserver sslmode=disable")
	if err != nil {
		log.Println("Error connecting to the database")
		log.Fatalln(err)
		panic(err)
	}

	return conn
}
