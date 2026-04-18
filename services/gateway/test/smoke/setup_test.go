package smoke

import (
	"net/http"
	"testing"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runSetupSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	t.Run(
		"00 health", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, "/health", "", nil, nil)
			env.RequireStatus(status, http.StatusOK, body)
		},
	)

	t.Run(
		"01 register owner", func(t *testing.T) {
			var resp dto.TokenResponse
			status, body := env.RequestJSON(
				http.MethodPost,
				"/api/v1/auth/register",
				"",
				dto.RegisterRequest{Email: env.Data.OwnerEmail, Password: env.Data.Password},
				&resp,
			)
			env.RequireStatus(status, http.StatusOK, body)
			if resp.Access == "" || resp.Refresh == "" {
				t.Fatalf("register token response is invalid: %+v", resp)
			}
			env.Data.OwnerAccess = resp.Access
			env.Data.OwnerRefresh = resp.Refresh
		},
	)

	t.Run(
		"02 login admin", func(t *testing.T) {
			var resp dto.TokenResponse
			status, body := env.RequestJSON(
				http.MethodPost,
				"/api/v1/auth/login",
				"",
				dto.LoginRequest{Email: env.Data.AdminEmail, Password: env.Data.AdminPassword},
				&resp,
			)
			env.RequireStatus(status, http.StatusOK, body)
			if resp.Access == "" || resp.Refresh == "" {
				t.Fatalf("admin login token response is invalid: %+v", resp)
			}
			env.Data.AdminAccess = resp.Access
		},
	)

	t.Run(
		"03 login owner", func(t *testing.T) {
			var resp dto.TokenResponse
			status, body := env.RequestJSON(
				http.MethodPost,
				"/api/v1/auth/login",
				"",
				dto.LoginRequest{Email: env.Data.OwnerEmail, Password: env.Data.Password},
				&resp,
			)
			env.RequireStatus(status, http.StatusOK, body)
			if resp.Access == "" || resp.Refresh == "" {
				t.Fatalf("owner login token response is invalid: %+v", resp)
			}
			env.Data.OwnerAccess = resp.Access
		},
	)

	t.Run(
		"04 refresh owner token", func(t *testing.T) {
			var resp dto.TokenResponse
			status, body := env.RequestJSON(
				http.MethodPost,
				"/api/v1/auth/refresh",
				"",
				dto.RefreshTokenRequest{RefreshToken: env.Data.OwnerRefresh},
				&resp,
			)
			env.RequireStatus(status, http.StatusOK, body)
			if resp.Access == "" || resp.Refresh == "" {
				t.Fatalf("refresh token response is invalid: %+v", resp)
			}
		},
	)

	t.Run(
		"05 resolve owner id via users list", func(t *testing.T) {
			var resp dto.UsersResponse
			status, body := env.RequestJSON(
				http.MethodGet,
				"/api/v1/auth/users?page=1&limit=100",
				env.Data.AdminAccess,
				nil,
				&resp,
			)
			env.RequireStatus(status, http.StatusOK, body)

			for _, user := range resp.Users {
				if user.Email == env.Data.OwnerEmail {
					env.Data.OwnerID = user.Id
					break
				}
			}
			if env.Data.OwnerID == 0 {
				t.Fatalf("owner not found in users list")
			}
		},
	)
}
