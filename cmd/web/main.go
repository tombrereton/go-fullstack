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
		log.Default().Println("No .env file found, using environment variables")
	}

	var cfg server.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Default().Printf("Is Development mode enabled: %t", cfg.IsDevelopment)

	server := server.NewWebServer(&cfg)
	address := "localhost:" + cfg.Port
	log.Default().Printf("Listening at http://%s", address)
	http.ListenAndServe(":"+cfg.Port, server.Router)
}
