package handler

import (
	"auth/internal/http/dto/request"
	"auth/internal/http/dto/response"
	"auth/internal/repository/postgres/models"
	"fukuro-reserve/pkg/utils/jwt"
)

type Service interface {
	UserService
	TokenService
}

type Handler struct {
	svc        Service
	tokenCreds *jwt.TokenCredentials
}

func New(svc Service, token *jwt.TokenCredentials) *Handler {
	return &Handler{
		svc:        svc,
		tokenCreds: token,
	}
}

func (h *Handler) UserCreateRequestToEntity(req *request.UserCreate, hash string) *models.UserCreate {
	return &models.UserCreate{
		Username: req.Username,
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
