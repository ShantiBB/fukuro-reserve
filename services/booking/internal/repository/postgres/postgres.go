package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"booking/internal/config"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(config *config.Config) (*Repository, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.DB,
		config.Postgres.SSLMode,
	)

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = 20
	cfg.MinConns = 5

	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.HealthCheckPeriod = 1 * time.Minute

	cfg.ConnConfig.ConnectTimeout = 10 * time.Second

	db, err := pgxpool.NewWithConfig(context.Background(), cfg)

	return &Repository{db: db}, nil
}
