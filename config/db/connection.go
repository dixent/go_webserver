package db

import (
	"log"
	"os"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Connection *sqlx.DB

func InitConnection() *sqlx.DB {
	log.Println(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	connParamsString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	conn, err := sqlx.Connect(os.Getenv("DB_ADAPTER"), connParamsString)

	if err != nil {
		log.Println("Error connecting to the database")
		log.Fatalln(err)
		panic(err)
	}

	return conn
}
