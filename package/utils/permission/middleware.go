package permission

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"auth_service/package/utils/helper"
)

type contextKey string
type Check func(r *http.Request, claims *Claims) bool

const ClaimsKey contextKey = "claims"

func AuthRequired(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				helper.SendError(w, r, http.StatusUnauthorized, ErrorResp("unauthorized"))
				return
			}

			claims, err := parseJWT(token, jwtSecret)
			if err != nil {
				helper.SendError(w, r, http.StatusForbidden, ErrorResp("forbidden"))
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetClaims(ctx context.Context) *Claims {
	claims, ok := ctx.Value(ClaimsKey).(*Claims)
	if !ok {
		return nil
	}
	return claims
}

func Require(checks ...Check) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := GetClaims(r.Context())
			if claims == nil {
				helper.SendError(w, r, http.StatusUnauthorized, ErrorResp("unauthorized"))
				return
			}

			for _, check := range checks {
				if check(r, claims) {
					next.ServeHTTP(w, r)
					return
				}
			}

			helper.SendError(w, r, http.StatusForbidden, ErrorResp("forbidden"))
		})
	}
}

func IsOwner(r *http.Request, claims *Claims) bool {
	userID := chi.URLParam(r, "id")
	targetUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return false
	}
	return claims.Sub == targetUserID
}

func IsAdmin(_ *http.Request, claims *Claims) bool {
	return claims.Role == RoleAdmin
}

func IsModerator(_ *http.Request, claims *Claims) bool {
	return claims.Role == RoleModerator
}
