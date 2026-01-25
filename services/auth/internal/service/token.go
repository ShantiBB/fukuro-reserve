package service

import (
	"context"
	"errors"
	"log/slog"

	"auth/internal/grpc/lib/utils/helper"
	"auth/internal/repository/models"
	"auth/pkg/lib/utils/consts"
	"auth/pkg/lib/utils/jwt"
)

func (s *Service) RegisterByEmail(ctx context.Context, user *models.CreateUser) (*jwt.Token, error) {
	created, err := s.repo.InsertUser(ctx, user)
	if err != nil {
		slog.Error("failed create user", "err:", err.Error())
		return nil, err
	}

	return jwt.GenerateAllTokens(created.ID, created.Role, s.tokenCreds)
}

func (s *Service) LoginByEmail(ctx context.Context, user *models.CreateUser) (*jwt.Token, error) {
	userCred, err := s.repo.SelectUserCredentialsByEmail(ctx, user.Email)
	if err != nil {
		slog.Error("failed login user", "err:", err.Error())
		return nil, err
	}

	if !helper.VerifyPassword(user.Password, userCred.Password) {
		return nil, consts.ErrInvalidCredentials
	}

	return jwt.GenerateAllTokens(userCred.ID, userCred.Role, s.tokenCreds)
}

func (s *Service) RefreshToken(token *jwt.Token) (*jwt.Token, error) {
	claims, err := jwt.GetClaimsRefreshToken(token.Refresh, s.tokenCreds.RefreshSecret)
	if err != nil {
		if errors.Is(err, consts.ErrInvalidToken) {
			return nil, consts.ErrInvalidToken
		}
		return nil, err
	}

	access, err := jwt.GenerateAccessToken(claims.Sub, claims.Role, s.tokenCreds)
	if err != nil {
		return nil, err
	}

	token.Access = access
	return token, nil
}
