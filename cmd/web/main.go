package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/tombrereton/go-hot-reload/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	var cfg server.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := server.NewWebServer(&cfg)
	address := "localhost:3000"
	log.Default().Printf("Listening at http://%s", address)
	http.ListenAndServe(address, server.Router)
}
