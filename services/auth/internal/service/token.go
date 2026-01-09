package service

import (
	"context"
	"errors"
	"log/slog"

	"auth/internal/repository/postgres/models"
	"auth/pkg/utils/consts"
	"auth/pkg/utils/jwt"
	"auth/pkg/utils/password"
)

func (s *Service) RegisterByEmail(ctx context.Context, email, password string) (*jwt.Token, error) {
	newUser := models.UserCreate{
		Email:    email,
		Password: password,
	}

	user, err := s.repo.UserCreate(ctx, newUser)
	if err != nil {
		slog.Error("failed create user", "err:", err.Error())
		return nil, err
	}

	return jwt.GenerateAllTokens(user.ID, user.Role, s.tokenCreds)
}

func (s *Service) LoginByEmail(ctx context.Context, email, pass string) (*jwt.Token, error) {
	user, err := s.repo.UserGetCredentialsByEmail(ctx, email)
	if err != nil {
		slog.Error("failed login user", "err:", err.Error())
		return nil, err
	}

	if !password.VerifyPassword(pass, user.Password) {
		return nil, consts.InvalidCredentials
	}

	return jwt.GenerateAllTokens(user.ID, user.Role, s.tokenCreds)
}

func (s *Service) RefreshToken(token *jwt.Token) (*jwt.Token, error) {
	claims, err := jwt.GetClaimsRefreshToken(s.tokenCreds.RefreshSecret, token.Refresh)
	if err != nil {
		if errors.Is(err, consts.InvalidRefreshToken) {
			return nil, consts.InvalidRefreshToken
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
