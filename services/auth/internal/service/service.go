package service

import (
	"auth/internal/database/postgres"
	"fukuro-reserve/pkg/utils/jwt"
)

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
