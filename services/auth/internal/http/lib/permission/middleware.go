package permission

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"auth/internal/http/dto/response"
	"auth/internal/http/lib/helper"
	"fukuro-reserve/pkg/utils/errs"
	"fukuro-reserve/pkg/utils/jwt"
)

type contextKey string
type Check func(r *http.Request, claims *jwt.Claims) bool

const ClaimsKey contextKey = "claims"

func AuthRequire(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				helper.SendError(w, r, http.StatusUnauthorized, response.ErrorResp(errs.Unauthorized))
				return
			}

			claims, err := jwt.ParseToken(token, []byte(jwtSecret))
			if err != nil {
				helper.SendError(w, r, http.StatusForbidden, response.ErrorResp(errs.Forbidden))
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetClaims(ctx context.Context) *jwt.Claims {
	claims, ok := ctx.Value(ClaimsKey).(*jwt.Claims)
	if !ok {
		return nil
	}
	return claims
}

func RequireRoles(checks ...Check) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := GetClaims(r.Context())
			if claims == nil {
				helper.SendError(w, r, http.StatusUnauthorized, response.ErrorResp(errs.Unauthorized))
				return
			}

			for _, check := range checks {
				if check(r, claims) {
					next.ServeHTTP(w, r)
					return
				}
			}

			helper.SendError(w, r, http.StatusForbidden, response.ErrorResp(errs.Forbidden))
		})
	}
}

func IsOwner(r *http.Request, claims *jwt.Claims) bool {
	userID := chi.URLParam(r, "id")
	targetUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return false
	}
	return claims.Sub == targetUserID
}

func IsAdmin(_ *http.Request, claims *jwt.Claims) bool {
	return claims.Role == RoleAdmin
}

func IsModerator(_ *http.Request, claims *jwt.Claims) bool {
	return claims.Role == RoleModerator
}
