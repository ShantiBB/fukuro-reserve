package service

import (
	"context"

	"auth/internal/repository/models"
)

type UserRepository interface {
	UserCreate(ctx context.Context, user models.UserCreate) (*models.User, error)
	UserGetByID(ctx context.Context, id int64) (*models.User, error)
	UserGetCredentialsByEmail(ctx context.Context, email string) (*models.UserCredentials, error)
	UserGetAll(ctx context.Context, limit, offset uint64) (*models.UserList, error)
	UserUpdateByID(ctx context.Context, user *models.User) error
	UserUpdateRoleStatus(ctx context.Context, id int64, role string) error
	UserUpdateActiveStatus(ctx context.Context, id int64, status bool) error
	UserDeleteByID(ctx context.Context, id int64) error
}

func (s *Service) UserCreate(ctx context.Context, user models.UserCreate) (*models.User, error) {
	newUser, err := s.repo.UserCreate(ctx, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) UserGetAll(ctx context.Context, page, limit uint64) (*models.UserList, error) {
	offset := (page - 1) * limit
	userList, err := s.repo.UserGetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return userList, nil
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

func (s *Service) UserUpdateRoleStatus(ctx context.Context, id int64, role string) error {
	if err := s.repo.UserUpdateRoleStatus(ctx, id, role); err != nil {
		return err
	}

	return nil
}

func (s *Service) UserUpdateActiveStatus(ctx context.Context, id int64, status bool) error {
	if err := s.repo.UserUpdateActiveStatus(ctx, id, status); err != nil {
		return err
	}

	return nil
}

func (s *Service) UserDeleteByID(ctx context.Context, id int64) error {
	err := s.repo.UserDeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
