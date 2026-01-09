package unit

import (
	"time"

	"auth/internal/http/dto/request"
	"auth/internal/repository/postgres/models"
	"auth/pkg/utils/jwt"
)

var (
	UserReq = request.UserCreate{
		Email:    "test@example.com",
		Password: "password123",
	}

	usernameMock = "test-user"
	UserMock     = models.User{
		ID:        1,
		Email:     "test@example.com",
		Username:  &usernameMock,
		Role:      "user",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	usernameShortMock = "test-user-short"
	UserShortMock     = models.UserShort{
		ID:       2,
		Email:    "testshort@example.com",
		Username: &usernameShortMock,
		Role:     "userShort",
		IsActive: true,
	}

	RegisterReq = request.UserCreate{
		Email:    "test@example.com",
		Password: "password123",
	}

	LoginReq = request.UserCreate{
		Email:    "test@example.com",
		Password: "password123",
	}

	RefreshReq = jwt.RefreshToken{
		RefreshToken: "valid-refresh-tokenCreds",
	}

	TokensMock = jwt.Token{
		Access:  "access-tokenCreds",
		Refresh: "refresh-tokenCreds",
	}

	LoginBadEmailAndPasswordReq = request.UserCreate{
		Email:    "test.com",
		Password: "123",
	}
)
