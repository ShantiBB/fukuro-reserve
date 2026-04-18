package auth

import (
	"context"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
)

func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.TokenResponse, error) {
	resp, err := s.clients.Token.RegisterUser(ctx, &userv1.RegisterUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return mapper.TokenResponseFromRegister(resp), nil
}

func (s *Service) Login(ctx context.Context, req dto.LoginRequest) (*dto.TokenResponse, error) {
	resp, err := s.clients.Token.LoginUser(ctx, &userv1.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return mapper.TokenResponseFromLogin(resp), nil
}

func (s *Service) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	resp, err := s.clients.Token.RefreshToken(ctx, &userv1.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	return mapper.TokenResponseFromRefresh(resp), nil
}
