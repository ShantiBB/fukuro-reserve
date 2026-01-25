package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"auth/pkg/lib/utils/consts"
)

func ParseBearerToken(tokenStr string, secret string) (*Claims, error) {
	tokenFunc := func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	}
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, tokenFunc)
	if err != nil {
		return nil, consts.ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func GetClaimsRefreshToken(tokenStr, refreshSecret string) (*Claims, error) {
	return ParseBearerToken(tokenStr, refreshSecret)
}
