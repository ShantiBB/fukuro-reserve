package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(sdn string) (*Repository, error) {
	db, err := sqlx.Connect("postgres", sdn)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}
