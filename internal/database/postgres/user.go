package postgres

import (
	"context"
	"fmt"

	"auth_service/internal/domain/models"
)

func (p *Repository) UserCreate(ctx context.Context, u models.UserCreate) (*models.User, error) {
	newUser := u.ToUserRead()
	err := p.db.QueryRow(
		ctx, UserCreate, u.Username, u.Email, u.Password,
	).Scan(&newUser.ID, &newUser.Role, &newUser.IsActive, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (p *Repository) UserGetByID(ctx context.Context, id int64) (*models.User, error) {
	u := models.User{ID: id}
	if err := p.db.QueryRow(ctx, UserGetByID, id).Scan(
		&u.Username, &u.Email, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *Repository) UserGetByUsername(ctx context.Context, username string) (*models.User, error) {
	u := models.User{Username: username}
	if err := p.db.QueryRow(ctx, UserGetByUsername, username).Scan(
		&u.ID, &u.Email, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *Repository) UserGetByEmail(ctx context.Context, email string) (*models.User, error) {
	u := models.User{Email: email}
	if err := p.db.QueryRow(ctx, UserGetByEmail, email).Scan(
		&u.ID, &u.Username, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *Repository) UserList(ctx context.Context) ([]models.User, error) {
	var users []models.User

	rows, err := p.db.Query(ctx, UserGetAll)
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

func (p *Repository) UserUpdateByID(ctx context.Context, u *models.User) error {
	rows, err := p.db.Exec(
		ctx, UserUpdate, u.Username, u.Email, u.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected := rows.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d is not found", u.ID)
	}

	return nil
}

func (p *Repository) UserDeleteByID(ctx context.Context, id int64) error {
	rows, err := p.db.Exec(ctx, UserDelete, id)
	if err != nil {
		return err
	}

	rowsAffected := rows.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d is not found", id)
	}

	return nil
}
