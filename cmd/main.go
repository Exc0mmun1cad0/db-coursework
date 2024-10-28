package main

import (
	"db-coursework/internal/config"
	"db-coursework/pkg/postgresql"
	"log"
)

func main() {
	cfg := config.MustLoad()

	conn, err := postgresql.NewClient(cfg.PostgreSQL)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}
	_ = conn

	postgresql.Migrate(cfg.PostgreSQL)

}
