package service

import (
	"auth/internal/repository/postgres"
	"auth/pkg/utils/jwt"
)

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
