package main

import (
	"context"
	"log"

	"my_vocabs/internal/my_vocabs/apis"
	"my_vocabs/internal/my_vocabs/config"
	db "my_vocabs/internal/my_vocabs/db/sqlc"
	fs "my_vocabs/pkg/file_store"

	"github.com/jackc/pgx/v5"
)

func main() {
	cfg, err := config.LoadConfigAndSecrets()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.Secrets.DBSource)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	// TODO: replace by log pkg
	log.Printf("Connected to DB successfully!!!")
	dbStore := db.NewStore(conn)

	// Connect to file store (GCS)
	var fileStore fs.FileStore
	if cfg.Env != config.LOCAL {
		fileStore, err = fs.NewGCSFileStore(cfg.Secrets.GSAEmail, cfg.Secrets.GSAKey)
		if err != nil {
			log.Fatal("cannot connect to file store: ", err)
		}
	}

	runHTTPServer(cfg, dbStore, fileStore)
}

func runHTTPServer(config *config.Config, dbStore db.Store, fileStore fs.FileStore) {
	server, err := apis.NewServer(config, dbStore, fileStore)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
