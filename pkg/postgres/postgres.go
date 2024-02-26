package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	*sqlx.DB
}

func NewPostgres() *Postgres {
	// postgres
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

	postgres := Postgres{conn}

	return &postgres
}
