package handler

import (
	"time"

	"auth/internal/domain/models"
	"auth/internal/http/lib/schemas/request"
	"fukuro-reserve/pkg/utils/jwt"
)

var (
	userReq = request.UserCreate{
		Username: "test",
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

	refreshReq = jwt.RefreshToken{
		RefreshToken: "valid-refresh-tokenCreds",
	}

	tokensMock = jwt.Token{
		Access:  "access-tokenCreds",
		Refresh: "refresh-tokenCreds",
	}

	loginBadEmailAndPasswordReq = request.UserCreate{
		Email:    "test.com",
		Password: "123",
	}
)
