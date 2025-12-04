package service

import (
	"context"

	"auth/internal/domain/models"
)

type UserRepository interface {
	UserCreate(ctx context.Context, user models.UserCreate) (*models.User, error)
	UserGetByID(ctx context.Context, id int64) (*models.User, error)
	UserGetCredentialsByEmail(ctx context.Context, email string) (*models.UserCredentials, error)
	UserList(ctx context.Context) ([]models.User, error)
	UserUpdateByID(ctx context.Context, user *models.User) error
	UserDeleteByID(ctx context.Context, id int64) error
}

func (s *Service) UserCreate(ctx context.Context, user models.UserCreate) (*models.User, error) {
	newUser, err := s.repo.UserCreate(ctx, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) UserList(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.UserList(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) UserGetByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.repo.UserGetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UserUpdateByID(ctx context.Context, user *models.User) (*models.User, error) {
	if err := s.repo.UserUpdateByID(ctx, user); err != nil {
		return nil, err
	}

	updatedUser, err := s.repo.UserGetByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *Service) UserDeleteByID(ctx context.Context, id int64) error {
	err := s.repo.UserDeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
