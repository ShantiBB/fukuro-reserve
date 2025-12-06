package service

import (
	"auth/internal/repository/postgres"
	"fukuro-reserve/pkg/utils/jwt"
)

type Repository interface {
	UserRepository
}

type Service struct {
	repo       *postgres.Repository
	tokenCreds *jwt.TokenCredentials
}

func New(repo *postgres.Repository, token *jwt.TokenCredentials) *Service {
	return &Service{
		repo:       repo,
		tokenCreds: token,
	}
}
