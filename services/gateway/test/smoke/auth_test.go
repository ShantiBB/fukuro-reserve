package smoke

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runAuthSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	t.Run("10 users list as owner forbidden", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/auth/users?page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			nil,
		)
		env.RequireStatus(status, http.StatusForbidden, body)
	})

	t.Run("11 users list as admin", func(t *testing.T) {
		var resp dto.UsersResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/auth/users?page=1&limit=10",
			env.Data.AdminAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("12 get owner profile", func(t *testing.T) {
		var resp dto.UserResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/auth/users/"+strconv.FormatInt(env.Data.OwnerID, 10),
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("13 update owner profile", func(t *testing.T) {
		var resp dto.UserResponse
		status, body := env.RequestJSON(
			http.MethodPut,
			"/api/v1/auth/users/"+strconv.FormatInt(env.Data.OwnerID, 10),
			env.Data.OwnerAccess,
			dto.UpdateUserRequest{Email: env.Data.OwnerEmail, Username: env.Data.OwnerUsername},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("14 create managed user as admin", func(t *testing.T) {
		var resp dto.UserResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/auth/users",
			env.Data.AdminAccess,
			dto.CreateUserRequest{Email: env.Data.ManagedEmail, Username: env.Data.ManagedUsername, Password: env.Data.Password},
			&resp,
		)
		env.RequireStatus(status, http.StatusCreated, body)
		env.Data.ManagedID = resp.Id
	})

	t.Run("15 update owner activity as admin", func(t *testing.T) {
		var resp dto.UpdateUserActivityResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/auth/users/"+strconv.FormatInt(env.Data.OwnerID, 10)+"/activity",
			env.Data.AdminAccess,
			dto.UpdateUserActivityRequest{IsActive: true},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("16 update owner role as admin", func(t *testing.T) {
		var resp dto.UpdateUserRoleResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/auth/users/"+strconv.FormatInt(env.Data.OwnerID, 10)+"/role",
			env.Data.AdminAccess,
			dto.UpdateUserRoleRequest{Role: "USER_ROLE_MODERATOR"},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})
}
