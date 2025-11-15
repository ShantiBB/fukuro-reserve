package handler

import (
	"time"

	"auth_service/internal/domain/models"
	"auth_service/internal/http/lib/schemas/request"
	"auth_service/package/utils/jwt"
)

var (
	userReq = request.UserCreate{
		Email:    "test@example.com",
		Password: "password123",
	}

	userMock = models.User{
		ID:        1,
		Email:     "test@example.com",
		Username:  "test-user",
		Role:      "user",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	registerReq = request.UserCreate{
		Email:    "test@example.com",
		Password: "password123",
	}

	loginReq = request.UserCreate{
		Email:    "test@example.com",
		Password: "password123",
	}

	refreshReq = request.RefreshToken{
		RefreshToken: "valid-refresh-token",
	}

	tokensMock = jwt.Token{
		Access:  "access-token",
		Refresh: "refresh-token",
	}

	loginBadEmailAndPasswordReq = request.UserCreate{
		Email:    "test.com",
		Password: "123",
	}
)
