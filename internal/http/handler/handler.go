package handler

import (
	"auth_service/internal/config"
	"auth_service/internal/domain/models"
	"auth_service/internal/http/lib/schemas/request"
	"auth_service/internal/http/lib/schemas/response"
)

type Service interface {
	UserService
	TokenService
}

type Handler struct {
	svc Service
	cfg *config.Config
}

func New(svc Service, cfg *config.Config) *Handler {
	return &Handler{
		svc: svc,
		cfg: cfg,
	}
}

func (h *Handler) UserCreateRequestToEntity(req *request.UserCreate, hash string) *models.UserCreate {
	return &models.UserCreate{
		Email:    req.Email,
		Password: hash,
	}
}

func (h *Handler) UserUpdateRequestToEntity(req *request.UserUpdate, id int64) *models.User {
	return &models.User{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
	}
}

func (h *Handler) UserEntityToResponse(user *models.User) *response.User {
	return &response.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
