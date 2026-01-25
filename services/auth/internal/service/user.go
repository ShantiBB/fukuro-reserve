package service

import (
	"context"

	"auth/internal/repository/models"
)

func (s *Service) CreateUser(ctx context.Context, user *models.CreateUser) (*models.User, error) {
	newUser, err := s.repo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) GetUsers(ctx context.Context, page, limit uint64) (*models.UserList, error) {
	offset := (page - 1) * limit
	userList, err := s.repo.SelectUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.repo.SelectUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUserByID(ctx context.Context, user *models.UpdateUser) error {
	if err := s.repo.UpdateUserByID(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateUserRoleStatus(ctx context.Context, id int64, role models.UserRole) error {
	if err := s.repo.UpdateUserRoleStatus(ctx, id, role); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateUserActiveStatus(ctx context.Context, id int64, status bool) error {
	if err := s.repo.UpdateUserActiveStatus(ctx, id, status); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUserByID(ctx context.Context, id int64) error {
	err := s.repo.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
