package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth/internal/config"
	"auth/internal/repository/models"
)

type Claims struct {
	jwt.RegisteredClaims
	Role models.UserRole
	Sub  int64
}

type Token struct {
	Access  string
	Refresh string
}

type TokenCredentials struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func GetTokenCredentials(cfg *config.Config) *TokenCredentials {
	return &TokenCredentials{
		AccessSecret:  cfg.JWT.AccessSecret,
		RefreshSecret: cfg.JWT.RefreshSecret,
		AccessTTL:     cfg.JWT.AccessTTL,
		RefreshTTL:    cfg.JWT.RefreshTTL,
	}
}
