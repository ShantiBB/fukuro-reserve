package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(sub int64, role string, ttl time.Duration, secret []byte) (string, error) {
	claims := Claims{
		Sub:  sub,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func GenerateAccessToken(sub int64, role string, t *TokenCredentials) (string, error) {
	return GenerateToken(sub, role, t.AccessTTL, []byte(t.AccessSecret))
}

func GenerateRefreshToken(sub int64, role string, t *TokenCredentials) (string, error) {
	return GenerateToken(sub, role, t.RefreshTTL, []byte(t.RefreshSecret))
}

func GenerateAllTokens(sub int64, role string, t *TokenCredentials) (*Token, error) {
	var err error
	tokens := &Token{}
	tokens.Access, err = GenerateAccessToken(sub, role, t)
	if err != nil {
		return nil, err
	}

	tokens.Refresh, err = GenerateRefreshToken(sub, role, t)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
