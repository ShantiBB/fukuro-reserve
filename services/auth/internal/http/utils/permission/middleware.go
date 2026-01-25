package permission

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"auth/internal/http/dto/response"
	"auth/internal/http/utils/helper"
	"auth/pkg/lib/utils/consts"
	"auth/pkg/lib/utils/jwt"
)

type contextKey string
type Check func(r *http.Request, claims *jwt.Claims) bool

const ClaimsKey contextKey = "claims"

func AuthRequire(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				bearerToken := r.Header.Get("Authorization")
				token, ok := strings.CutPrefix(bearerToken, "Bearer ")
				if !ok {
					helper.SendError(w, r, http.StatusUnauthorized, response.ErrorResp(consts.ErrUnauthorized))
					return
				}

				claims, err := jwt.ParseBearerToken(token, jwtSecret)
				if err != nil {
					helper.SendError(w, r, http.StatusForbidden, response.ErrorResp(consts.ErrForbidden))
					return
				}

				ctx := context.WithValue(r.Context(), ClaimsKey, claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
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
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				claims := GetClaims(r.Context())
				if claims == nil {
					helper.SendError(w, r, http.StatusUnauthorized, response.ErrorResp(consts.ErrUnauthorized))
					return
				}

				for _, check := range checks {
					if check(r, claims) {
						next.ServeHTTP(w, r)
						return
					}
				}

				helper.SendError(w, r, http.StatusForbidden, response.ErrorResp(consts.ErrForbidden))
			},
		)
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
