package main

import (
	"flag"
	"fmt"
	"go_webserver/config"
	"go_webserver/internal/auth"
	"go_webserver/internal/shop"
	"go_webserver/pkg/postgres"
	"log"
	"net/http"
)

func main() {
	config.InitEnvironment()
	postgres := postgres.NewPostgres()

	port := flag.Int("port", 3000, "the http port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/api/shops/", shop.Handler{})
	mux.Handle("/api/shops", shop.Handler{})

	auth := auth.NewRunner(mux, postgres)
	auth.Run()

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), mux))

	defer postgres.Close()
}
