package main

import (
	"db-coursework/internal/config"
	postgresql "db-coursework/pkg/postgresql_client"
	"log"
)

func main() {
	cfg := config.MustLoad()

	conn, err := postgresql.NewClient(cfg.PostgreSQL)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	_ = conn
}
