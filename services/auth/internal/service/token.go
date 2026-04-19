package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/grpc/lib/utils/helper"
	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/repository/models"
	"github.com/ShantiBB/fukuro-reserve/services/auth/pkg/lib/utils/consts"
	"github.com/ShantiBB/fukuro-reserve/services/auth/pkg/lib/utils/jwt"
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
		if errors.Is(err, consts.ErrUserNotFound) {
			return nil, consts.ErrInvalidCredentials
		}
		return nil, err
	}
	if !userCred.IsActive {
		return nil, consts.ErrForbidden
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
	user, err := s.repo.SelectUserByID(context.Background(), claims.Sub)
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, consts.ErrForbidden
	}

	access, err := jwt.GenerateAccessToken(claims.Sub, claims.Role, s.tokenCreds)
	if err != nil {
		return nil, err
	}

	token.Access = access
	return token, nil
}
