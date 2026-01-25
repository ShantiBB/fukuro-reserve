package handler

import (
	"auth/internal/http/dto/request"
	"auth/internal/http/dto/response"
	"auth/internal/repository/models"
)

type Service interface {
	UserService
	TokenService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) UserCreateRequestToEntity(req *request.UserCreate, hash string) *models.CreateUser {
	return &models.CreateUser{
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

func (h *Handler) UserShortEntityToResponse(user *models.UserShort) *response.UserShort {
	return &response.UserShort{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		IsActive: user.IsActive,
	}
}
