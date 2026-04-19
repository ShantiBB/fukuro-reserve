package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/repository/models"
)

type Claims struct {
	jwt.RegisteredClaims
	Role models.UserRole `json:"role"`
	Sub  int64           `json:"sub"`
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
