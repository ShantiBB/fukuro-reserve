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
	validRoles := map[string]struct{}{
		"USER_ROLE_USER":      {},
		"USER_ROLE_MODERATOR": {},
		"USER_ROLE_ADMIN":     {},
	}

	t.Run("10 users list as owner forbidden", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/users?page=1&limit=10",
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
			"/api/v1/users?page=1&limit=10",
			env.Data.AdminAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if len(resp.Users) == 0 {
			t.Fatalf("users list is empty")
		}
		for _, u := range resp.Users {
			if u.Id <= 0 || u.Email == "" {
				t.Fatalf("invalid user item in users list: %+v", u)
			}
			if _, ok := validRoles[u.Role]; !ok {
				t.Fatalf("unexpected role in users list: %q", u.Role)
			}
		}
	})

	t.Run("12 get owner profile", func(t *testing.T) {
		var resp dto.UserResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/users/"+strconv.FormatInt(env.Data.OwnerID, 10),
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Id != env.Data.OwnerID || resp.Email == "" {
			t.Fatalf("invalid owner profile response: %+v", resp)
		}
		if _, ok := validRoles[resp.Role]; !ok {
			t.Fatalf("unexpected owner role: %q", resp.Role)
		}
	})

	t.Run("13 update owner profile", func(t *testing.T) {
		var resp dto.UserResponse
		status, body := env.RequestJSON(
			http.MethodPut,
			"/api/v1/users/"+strconv.FormatInt(env.Data.OwnerID, 10),
			env.Data.OwnerAccess,
			dto.UpdateUserRequest{Email: env.Data.OwnerEmail, Username: env.Data.OwnerUsername},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Email != env.Data.OwnerEmail || resp.Username != env.Data.OwnerUsername {
			t.Fatalf("invalid updated owner response: %+v", resp)
		}
	})

	t.Run("14 create managed user as admin", func(t *testing.T) {
		var resp dto.UserResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/users",
			env.Data.AdminAccess,
			dto.CreateUserRequest{Email: env.Data.ManagedEmail, Username: env.Data.ManagedUsername, Password: env.Data.Password},
			&resp,
		)
		env.RequireStatus(status, http.StatusCreated, body)
		if resp.Id <= 0 || resp.Email != env.Data.ManagedEmail {
			t.Fatalf("invalid create managed user response: %+v", resp)
		}
		if _, ok := validRoles[resp.Role]; !ok {
			t.Fatalf("unexpected managed user role: %q", resp.Role)
		}
		env.Data.ManagedID = resp.Id
	})

	t.Run("14.1 login managed user", func(t *testing.T) {
		var resp dto.TokenResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/auth/login",
			"",
			dto.LoginRequest{Email: env.Data.ManagedEmail, Password: env.Data.Password},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Access == "" {
			t.Fatalf("managed login token response is invalid: %+v", resp)
		}
		env.Data.ManagedAccess = resp.Access
	})

	t.Run("15 update owner activity as admin", func(t *testing.T) {
		var resp dto.UpdateUserActivityResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/users/"+strconv.FormatInt(env.Data.OwnerID, 10)+"/activity",
			env.Data.AdminAccess,
			dto.UpdateUserActivityRequest{IsActive: true},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if !resp.IsActive {
			t.Fatalf("expected is_active=true, got %+v", resp)
		}
	})

	t.Run("16 update owner role as admin", func(t *testing.T) {
		var resp dto.UpdateUserRoleResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/users/"+strconv.FormatInt(env.Data.OwnerID, 10)+"/role",
			env.Data.AdminAccess,
			dto.UpdateUserRoleRequest{Role: "USER_ROLE_MODERATOR"},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Role != "USER_ROLE_MODERATOR" {
			t.Fatalf("unexpected role update response: %+v", resp)
		}
	})

	t.Run("17 set role user", func(t *testing.T) {
		var resp dto.UpdateUserRoleResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/users/"+strconv.FormatInt(env.Data.OwnerID, 10)+"/role",
			env.Data.AdminAccess,
			dto.UpdateUserRoleRequest{Role: "USER_ROLE_USER"},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Role != "USER_ROLE_USER" {
			t.Fatalf("unexpected role update response: %+v", resp)
		}
	})

	t.Run("18 set role admin", func(t *testing.T) {
		var resp dto.UpdateUserRoleResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/users/"+strconv.FormatInt(env.Data.OwnerID, 10)+"/role",
			env.Data.AdminAccess,
			dto.UpdateUserRoleRequest{Role: "USER_ROLE_ADMIN"},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Role != "USER_ROLE_ADMIN" {
			t.Fatalf("unexpected role update response: %+v", resp)
		}
	})

	t.Run("19 set role moderator for next tests", func(t *testing.T) {
		var resp dto.UpdateUserRoleResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/users/"+strconv.FormatInt(env.Data.OwnerID, 10)+"/role",
			env.Data.AdminAccess,
			dto.UpdateUserRoleRequest{Role: "USER_ROLE_MODERATOR"},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Role != "USER_ROLE_MODERATOR" {
			t.Fatalf("unexpected role update response: %+v", resp)
		}
	})

	t.Run("19.1 login owner after role change", func(t *testing.T) {
		var resp dto.TokenResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/auth/login",
			"",
			dto.LoginRequest{Email: env.Data.OwnerEmail, Password: env.Data.Password},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Access == "" {
			t.Fatalf("owner re-login token response is invalid: %+v", resp)
		}
		env.Data.OwnerAccess = resp.Access
	})
}
