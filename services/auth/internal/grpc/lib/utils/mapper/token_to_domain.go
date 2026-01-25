package mapper

import (
	userv1 "auth/api/user/v1"
	"auth/pkg/lib/utils/jwt"
)

func RefreshTokenRequestToDomain(req *userv1.RefreshTokenRequest) *jwt.Token {
	return &jwt.Token{
		Refresh: req.RefreshToken,
	}
}
