package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	jwtclaims "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/jwt"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
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
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(consts.HeaderAuthorization)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: "authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: "invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: "invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: "invalid token claims"})
			c.Abort()
			return
		}

		userID, ok := jwtclaims.ExtractUserID(claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: "invalid user id in token"})
			c.Abort()
			return
		}

		userEmail := jwtclaims.ExtractStringClaim(claims, "email", "Email")
		userRole := jwtclaims.ExtractStringClaim(claims, "role", "Role")

		ctx := context.WithValue(c.Request.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, UserEmailKey, userEmail)
		ctx = context.WithValue(ctx, UserRoleKey, userRole)
		c.Request = c.Request.WithContext(ctx)

		c.Set(string(UserIDKey), userID)
		c.Set(string(UserEmailKey), userEmail)
		c.Set(string(UserRoleKey), userRole)

		c.Next()
	}
}
