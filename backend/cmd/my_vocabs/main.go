package main

import (
	"log"

	"my_vocabs/internal/my_vocabs/apis"
	"my_vocabs/internal/my_vocabs/config"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	runHTTPServer(cfg)
}

func runHTTPServer(config *config.Config) {
	server, err := apis.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
