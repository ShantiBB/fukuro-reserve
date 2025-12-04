package service

import (
	"context"
	"errors"

	"auth/internal/domain/models"
	"fukuro-reserve/pkg/utils/errs"
	"fukuro-reserve/pkg/utils/jwt"
	"fukuro-reserve/pkg/utils/password"
)

func (s *Service) RegisterByEmail(ctx context.Context, email, password string) (*jwt.Token, error) {
	newUser := models.UserCreate{
		Email:    email,
		Password: password,
	}

	user, err := s.repo.UserCreate(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return jwt.GenerateAllTokens(user.ID, user.Role, s.tokenCreds)
}

func (s *Service) LoginByEmail(ctx context.Context, email, pass string) (*jwt.Token, error) {
	user, err := s.repo.UserGetCredentialsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !password.VerifyPassword(pass, user.Password) {
		return nil, errs.InvalidCredentials
	}

	return jwt.GenerateAllTokens(user.ID, user.Role, s.tokenCreds)
}

func (s *Service) RefreshToken(token *jwt.Token) (*jwt.Token, error) {
	claims, err := jwt.GetClaimsRefreshToken(s.tokenCreds.RefreshSecret, token.Refresh)
	if err != nil {
		if errors.Is(err, errs.InvalidRefreshToken) {
			return nil, errs.InvalidRefreshToken
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
