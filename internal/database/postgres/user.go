package postgres

import (
	"context"
	"fmt"

	"auth_service/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.UserCreate) (*models.User, error)
	Get(ctx context.Context, id int64) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id int64) error
}

func (p *Repository) Create(ctx context.Context, user models.UserCreate) (*models.User, error) {
	newUser := user.ToUserRead()
	rows, err := p.db.NamedQueryContext(ctx, UserCreate, user)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&newUser.ID); err != nil {
			return nil, err
		}
	}

	return newUser, nil
}

func (p *Repository) Get(ctx context.Context, id int64) (*models.User, error) {
	var user models.User

	if err := p.db.GetContext(ctx, &user, UserGet, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *Repository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User

	if err := p.db.SelectContext(ctx, &users, UserGetAll); err != nil {
		return nil, err
	}

	return users, nil
}

func (p *Repository) Update(ctx context.Context, user models.User) error {
	rows, err := p.db.NamedExecContext(ctx, UserUpdate, &user)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("user with id %d is not found", user.ID)
	}

	return nil
}

func (p *Repository) Delete(ctx context.Context, id int64) error {
	rows, err := p.db.ExecContext(ctx, UserDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d is not found", id)
	}

	return nil
}
