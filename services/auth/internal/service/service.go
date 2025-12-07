package service

import (
	"auth/internal/http/lib/jwt"
	"auth/internal/repository/postgres"
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
