package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"auth/internal/repository/models"
	"auth/pkg/lib/utils/consts"
)

func (r *Repository) InsertUser(ctx context.Context, u *models.CreateUser) (*models.User, error) {
	newUser := u.ToUserRead()
	err := r.db.QueryRow(ctx, InsertUser, u.Email, u.Username, u.Password).
		Scan(&newUser.ID, &newUser.Role, &newUser.IsActive, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, consts.ErrUniqueUserField
		}
		return nil, err
	}

	return newUser, nil
}

func (r *Repository) SelectUsers(ctx context.Context, limit, offset uint64) (*models.UserList, error) {
	rows, err := r.db.Query(ctx, SelectUsers, limit, offset)
	if err != nil {
		return nil, err
	}

	values := make([]models.UserShort, limit)
	var u models.UserShort
	var totalCount, idx uint64
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role, &u.IsActive, &totalCount)
		if err != nil {
			return nil, err
		}

		values[idx] = u
		idx++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	values = values[:idx]

	userList := &models.UserList{
		Users:      make([]*models.UserShort, len(values)),
		TotalCount: totalCount,
	}
	for i := range values {
		userList.Users[i] = &values[i]
	}

	return userList, nil
}

func (r *Repository) SelectUserByID(ctx context.Context, id int64) (*models.User, error) {
	u := &models.User{ID: id}
	if err := r.db.QueryRow(ctx, SelectUserByID, id).Scan(
		&u.Username, &u.Email, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, consts.ErrUserNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *Repository) SelectUserCredentialsByEmail(ctx context.Context, email string) (*models.UserCredentials, error) {
	u := &models.UserCredentials{Email: email}
	err := r.db.QueryRow(ctx, SelectUserCredentialsByEmail, email).Scan(
		&u.ID, &u.Role, &u.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, consts.ErrUserNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *Repository) UpdateUserByID(ctx context.Context, u *models.UpdateUser) error {
	rows, err := r.db.Exec(ctx, UpdateUser, u.ID, u.Username, u.Email)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.ErrUniqueUserField
		}
		return err
	}

	if rows.RowsAffected() == 0 {
		return consts.ErrUserNotFound
	}

	return nil
}

func (r *Repository) UpdateUserRoleStatus(ctx context.Context, id int64, role models.UserRole) error {
	row, err := r.db.Exec(ctx, UpdateUserRoleStatus, id, role)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "22P02" {
				return consts.ErrInvalidRole
			}
		}
		return err
	}

	if row.RowsAffected() == 0 {
		return consts.ErrUserNotFound
	}

	return nil
}

func (r *Repository) UpdateUserActiveStatus(ctx context.Context, id int64, status bool) error {
	row, err := r.db.Exec(ctx, UpdateUserActiveStatus, id, status)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return consts.ErrUserNotFound
	}

	return nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id int64) error {
	row, err := r.db.Exec(ctx, DeleteUser, id)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return consts.ErrUserNotFound
	}

	return nil
}
