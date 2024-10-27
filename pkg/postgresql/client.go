package postgresql

import (
	"db-coursework/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewClient(cfg config.PostgreSQL) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx failed to connect")
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping failed")
	}

	return conn, nil
}
