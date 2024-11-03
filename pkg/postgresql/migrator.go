package postgresql

import (
	"db-coursework/internal/config"
	"fmt"
	"log"
	"os"
	"time"

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
		err = m.Up()
		if err == migrate.ErrNoChange {
			log.Println("All migrations are applied")
			break
		} else if err != nil {
			return errors.Wrap(err, "Migrations failed")
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}
