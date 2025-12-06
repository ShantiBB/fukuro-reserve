package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"auth/internal/repository/postgres/models"
	"fukuro-reserve/pkg/utils/consts"
)

func (r *Repository) UserCreate(ctx context.Context, u models.UserCreate) (*models.User, error) {
	newUser := u.ToUserRead()
	err := r.db.QueryRow(
		ctx, UserCreate, u.Username, u.Email, u.Password,
	).Scan(&newUser.ID, &newUser.Role, &newUser.IsActive, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, consts.UniqueEmailField
		}
		return nil, err
	}

	return &newUser, nil
}

func (r *Repository) UserGetAll(ctx context.Context, limit, offset uint64) ([]models.User, error) {
	var users []models.User

	rows, err := r.db.Query(ctx, UserGetAll, limit, offset)
	if err != nil {
		return nil, err
	}

	var u models.User
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *Repository) UserGetByID(ctx context.Context, id int64) (*models.User, error) {
	u := models.User{ID: id}
	if err := r.db.QueryRow(ctx, UserGetByID, id).Scan(
		&u.Username, &u.Email, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, consts.UserNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *Repository) UserGetCredentialsByEmail(ctx context.Context, email string) (*models.UserCredentials, error) {
	u := models.UserCredentials{Email: email}
	err := r.db.QueryRow(ctx, UserGetCredentialsByEmail, email).Scan(
		&u.ID, &u.Role, &u.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, consts.UserNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *Repository) UserUpdateByID(ctx context.Context, u *models.User) error {
	rows, err := r.db.Exec(
		ctx, UserUpdate, u.Username, u.Email, u.ID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.UniqueEmailField
		}
		return err
	}

	rowsAffected := rows.RowsAffected()
	if rowsAffected == 0 {
		return consts.UserNotFound
	}

	return nil
}

func (r *Repository) UserDeleteByID(ctx context.Context, id int64) error {
	rows, err := r.db.Exec(ctx, UserDelete, id)
	if err != nil {
		return err
	}

	rowsAffected := rows.RowsAffected()
	if rowsAffected == 0 {
		return consts.UserNotFound
	}

	return nil
}

func (r *Repository) UserGetCountRows(ctx context.Context) (int, error) {
	var count int
	if err := r.db.QueryRow(ctx, UserGetCountRows).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
