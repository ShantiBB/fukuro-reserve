package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub  int64
	Role string
	jwt.RegisteredClaims
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
