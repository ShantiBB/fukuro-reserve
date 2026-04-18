package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/responder"
)

type contextKey string

const (
	UserIDKey    contextKey = "userID"
	UserRoleKey  contextKey = "userRole"
	UserEmailKey contextKey = "userEmail"
)

var jwtSecret []byte

// SetJWTSecret sets the JWT secret for token validation
func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}

// AuthMiddleware validates JWT tokens
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get(consts.HeaderAuthorization)
			if authHeader == "" {
				responder.Error(w, http.StatusUnauthorized, "authorization header is required")
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				responder.Error(w, http.StatusUnauthorized, "invalid authorization header format")
				return
			}

			tokenString := parts[1]
			token, err := jwt.Parse(
				tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, jwt.ErrSignatureInvalid
					}
					return jwtSecret, nil
				},
			)

			if err != nil || !token.Valid {
				responder.Error(w, http.StatusUnauthorized, "invalid token")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				responder.Error(w, http.StatusUnauthorized, "invalid token claims")
				return
			}

			userID, ok := extractUserID(claims)
			if !ok {
				responder.Error(w, http.StatusUnauthorized, "invalid user id in token")
				return
			}

			userEmail := extractStringClaim(claims, "email", "Email")
			userRole := extractStringClaim(claims, "role", "Role")

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			ctx = context.WithValue(ctx, UserEmailKey, userEmail)
			ctx = context.WithValue(ctx, UserRoleKey, userRole)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func extractUserID(claims jwt.MapClaims) (int64, bool) {
	for _, key := range []string{"sub", "Sub"} {
		switch value := claims[key].(type) {
		case float64:
			return int64(value), true
		case string:
			id, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				return id, true
			}
		}
	}

	return 0, false
}

func extractStringClaim(claims jwt.MapClaims, keys ...string) string {
	for _, key := range keys {
		if value, ok := claims[key].(string); ok {
			return value
		}
	}

	return ""
}
