package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"auth/internal/repository/postgres/models"
	"auth/pkg/utils/consts"
)

func (r *Repository) UserCreate(ctx context.Context, u models.UserCreate) (*models.User, error) {
	newUser := u.ToUserRead()
	err := r.db.QueryRow(ctx, UserCreate, u.Email, u.Password).
		Scan(&newUser.ID, &newUser.Role, &newUser.IsActive, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, consts.UniqueUserField
		}
		return nil, err
	}

	return &newUser, nil
}

func (r *Repository) UserGetAll(ctx context.Context, limit, offset uint64) (*models.UserList, error) {
	var userList models.UserList

	rows, err := r.db.Query(ctx, UserGetAll, limit, offset)
	if err != nil {
		return nil, err
	}

	var u models.UserShort
	for rows.Next() {
		if err = rows.
			Scan(&u.ID, &u.Username, &u.Email, &u.Role, &u.IsActive); err != nil {
			return nil, err
		}

		userList.Users = append(userList.Users, u)
	}

	if err = r.db.QueryRow(ctx, UserGetCountRows).Scan(&userList.TotalCount); err != nil {
		return nil, err
	}

	return &userList, nil
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
			return consts.UniqueUserField
		}
		return err
	}

	rowsAffected := rows.RowsAffected()
	if rowsAffected == 0 {
		return consts.UserNotFound
	}

	return nil
}

func (r *Repository) UserUpdateRoleStatus(ctx context.Context, id int64, role string) error {
	row, err := r.db.Exec(ctx, UserUpdateRoleStatus, role, id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "22P02" {
				return consts.ErrInvalidRole
			}
		}
		return err
	}

	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.UserNotFound
	}

	return nil
}

func (r *Repository) UserUpdateActiveStatus(ctx context.Context, id int64, status bool) error {
	row, err := r.db.Exec(ctx, UserUpdateActiveStatus, status, id)
	if err != nil {
		return err
	}

	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.UserNotFound
	}

	return nil
}

func (r *Repository) UserDeleteByID(ctx context.Context, id int64) error {
	row, err := r.db.Exec(ctx, UserDelete, id)
	if err != nil {
		return err
	}

	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.UserNotFound
	}

	return nil
}
