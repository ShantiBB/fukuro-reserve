package service

import (
	"context"

	"auth/internal/repository/models"
	"auth/internal/repository/postgres"
	"auth/pkg/lib/utils/jwt"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.CreateUser) (*models.User, error)
	SelectUserByID(ctx context.Context, id int64) (*models.User, error)
	SelectUserCredentialsByEmail(ctx context.Context, email string) (*models.UserCredentials, error)
	SelectUsers(ctx context.Context, limit, offset uint64) (*models.UserList, error)
	UpdateUserByID(ctx context.Context, user *models.UpdateUser) error
	UpdateUserRoleStatus(ctx context.Context, id int64, role models.UserRole) error
	UpdateUserActiveStatus(ctx context.Context, id int64, status bool) error
	DeleteUserByID(ctx context.Context, id int64) error
}

type Repository interface {
	UserRepository
}

type Service struct {
	repo       Repository
	tokenCreds *jwt.TokenCredentials
}

func New(repo *postgres.Repository, token *jwt.TokenCredentials) *Service {
	return &Service{
		repo:       repo,
		tokenCreds: token,
	}
}
