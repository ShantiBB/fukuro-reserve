package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"fukuro-reserve/pkg/utils/errs"
)

func ParseToken(tokenStr string, secret []byte) (*Claims, error) {
	tokenFunc := func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	}
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, tokenFunc)
	if err != nil {
		return nil, errs.InvalidRefreshToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func GetClaimsRefreshToken(refreshSecret, tokenStr string) (*Claims, error) {
	return ParseToken(tokenStr, []byte(refreshSecret))
}
