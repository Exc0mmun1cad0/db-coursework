package postgresql

import (
	"db-coursework/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

func Migrate(cfg config.PostgreSQL) error {
	var migrationsPath = os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		log.Fatal("MIGRATIONS_PATH is not set")
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	m, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), dbURL)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if err := m.Steps(1); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("All migrations applied")
				break
			}
			return errors.Wrap(err, "error during migration")
		} else {
			version, dirty, err := m.Version()
			if err != nil {
				return errors.Wrap(err, "error during migration")
			}
			log.Printf("Applied migration: %d, Dirty: %t\n", version, dirty)
		}
	}

	return nil
}
