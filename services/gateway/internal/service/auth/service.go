package auth

import (
	"context"

	"google.golang.org/protobuf/types/known/wrapperspb"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
)

type Service struct {
	clients    *clients.Clients
	pagination config.PaginationConfig
}

func New(clients *clients.Clients, pagination config.PaginationConfig) *Service {
	return &Service{clients: clients, pagination: pagination}
}

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

func (s *Service) GetUsers(ctx context.Context, page, limit uint64) (*dto.UsersResponse, error) {
	if page == 0 {
		page = s.pagination.DefaultPage
	}
	if limit == 0 {
		limit = s.pagination.DefaultPageSize
	}

	resp, err := s.clients.User.GetUsers(ctx, &userv1.GetUsersRequest{Page: page, Limit: limit})
	if err != nil {
		return nil, err
	}

	return mapper.UsersResponseFromProto(resp), nil
}

func (s *Service) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	var username *string
	if req.Username != "" {
		username = &req.Username
	}

	resp, err := s.clients.User.CreateUser(ctx, &userv1.CreateUserRequest{
		Email:    req.Email,
		Username: username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return mapper.UserResponseFromProto(resp.User), nil
}

func (s *Service) GetUser(ctx context.Context, id int64) (*dto.UserResponse, error) {
	resp, err := s.clients.User.GetUser(ctx, &userv1.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return mapper.UserResponseFromProto(resp.User), nil
}

func (s *Service) UpdateUser(ctx context.Context, id int64, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	resp, err := s.clients.User.UpdateUser(ctx, &userv1.UpdateUserRequest{
		Id:       id,
		Email:    req.Email,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	return mapper.UpdateUserResponseFromProto(resp.User), nil
}

func (s *Service) UpdateUserActivity(ctx context.Context, id int64, req dto.UpdateUserActivityRequest) (*dto.UpdateUserActivityResponse, error) {
	resp, err := s.clients.User.UpdateUserActivity(ctx, &userv1.UpdateUserActivityRequest{
		Id:       id,
		IsActive: wrapperspb.Bool(req.IsActive),
	})
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserActivityResponse{IsActive: resp.IsActive}, nil
}

func (s *Service) UpdateUserRole(ctx context.Context, id int64, req dto.UpdateUserRoleRequest) (*dto.UpdateUserRoleResponse, error) {
	resp, err := s.clients.User.UpdateUserRole(ctx, &userv1.UpdateUserRoleRequest{
		Id:   id,
		Role: userv1.UserRole(userv1.UserRole_value[req.Role]),
	})
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserRoleResponse{Role: resp.Role.String()}, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	_, err := s.clients.User.DeleteUser(ctx, &userv1.DeleteUserRequest{Id: id})
	return err
}
