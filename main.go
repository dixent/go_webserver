package main

import (
	"flag"
	"fmt"
	"go_webserver/db"
	"go_webserver/internal/shop"
	"log"
	"net/http"
)

func main() {
	db.Connection = db.InitConnection()

	port := flag.Int("port", 3000, "the http port")

	mux := http.NewServeMux()
	mux.Handle("/api/shops/", shop.Handler{})
	mux.Handle("/api/shops", shop.Handler{})

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), mux))

	defer db.Connection.Close()
}
