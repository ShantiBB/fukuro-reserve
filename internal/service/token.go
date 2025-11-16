package service

import (
	"context"
	"errors"

	"auth_service/internal/config"
	"auth_service/internal/domain/models"
	"auth_service/package/utils/errs"
	"auth_service/package/utils/jwt"
	"auth_service/package/utils/password"
)

func (s *Service) RegisterByEmail(
	ctx context.Context,
	email, password string,
	cfg *config.Config,
) (*jwt.Token, error) {
	newUser := models.UserCreate{
		Email:    email,
		Password: password,
	}

	user, err := s.repo.UserCreate(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return jwt.GenerateAllTokens(user.ID, user.Role, cfg)
}

func (s *Service) LoginByEmail(ctx context.Context, email, pass string, cfg *config.Config) (*jwt.Token, error) {
	user, err := s.repo.UserGetCredentialsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !password.VerifyPassword(pass, user.Password) {
		return nil, errs.InvalidCredentials
	}

	return jwt.GenerateAllTokens(user.ID, user.Role, cfg)
}

func (s *Service) RefreshToken(token *jwt.Token, cfg *config.Config) (*jwt.Token, error) {
	claims, err := jwt.GetClaimsRefreshToken(cfg.JWT.RefreshSecret, token.Refresh)
	if err != nil {
		if errors.Is(err, errs.InvalidRefreshToken) {
			return nil, errs.InvalidRefreshToken
		}
		return nil, err
	}

	access, err := jwt.GenerateAccessToken(claims.Sub, claims.Role, cfg.JWT.AccessSecret, cfg.JWT.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	token.Access = access
	return token, nil
}
