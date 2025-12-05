package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"hotel/internal/config"
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

	db, err := pgxpool.NewWithConfig(context.Background(), cfg)

	return &Repository{db: db}, nil
}
