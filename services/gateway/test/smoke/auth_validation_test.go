package smoke

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runAuthValidationSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	t.Run("auth grpc validation matrix", func(t *testing.T) {
		// Some auth gRPC proto validations are not reachable through gateway by design:
		// - id >= 1 checks for get/update/delete/activity/role (gateway path parser rejects first)
		// - get_users.page >= 1 and get_users.limit >= 1 (gateway/service apply defaults)
		// - update_user_activity.is_active required (gateway always sends BoolValue, even false)
		registerPath := "/api/v1/auth/register"
		loginPath := "/api/v1/auth/login"
		refreshPath := "/api/v1/auth/refresh"
		usersPath := "/api/v1/users"
		ownerIDPath := usersPath + "/" + strconv.FormatInt(env.Data.OwnerID, 10)

		t.Run("register email required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, registerPath, "", map[string]any{
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("register email format", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, registerPath, "", map[string]any{
				"email":    "invalid-email",
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("register email max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, registerPath, "", map[string]any{
				"email":    longEmail(96),
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("register password required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, registerPath, "", map[string]any{
				"email": "valid-register@example.com",
			}, "password")
		})
		t.Run("register password min length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, registerPath, "", map[string]any{
				"email":    "valid-register@example.com",
				"password": "1234567",
			}, "password")
		})
		t.Run("register password max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, registerPath, "", map[string]any{
				"email":    "valid-register@example.com",
				"password": strings.Repeat("p", 101),
			}, "password")
		})

		t.Run("login email required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, loginPath, "", map[string]any{
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("login email format", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, loginPath, "", map[string]any{
				"email":    "invalid-email",
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("login email max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, loginPath, "", map[string]any{
				"email":    longEmail(96),
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("login password required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, loginPath, "", map[string]any{
				"email": "valid-login@example.com",
			}, "password")
		})
		t.Run("login password min length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, loginPath, "", map[string]any{
				"email":    "valid-login@example.com",
				"password": "1234567",
			}, "password")
		})
		t.Run("login password max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, loginPath, "", map[string]any{
				"email":    "valid-login@example.com",
				"password": strings.Repeat("p", 101),
			}, "password")
		})

		t.Run("create user email required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("create user email format", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email":    "invalid-email",
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("create user email max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email":    longEmail(96),
				"password": env.Data.Password,
			}, "email")
		})
		t.Run("create user username max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email":    "create-user-username-max@example.com",
				"username": strings.Repeat("u", 101),
				"password": env.Data.Password,
			}, "username")
		})
		t.Run("create user password required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email": "create-user-password-required@example.com",
			}, "password")
		})
		t.Run("create user password min length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email":    "create-user-password-min@example.com",
				"password": "1234567",
			}, "password")
		})
		t.Run("create user password max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email":    "create-user-password-max@example.com",
				"password": strings.Repeat("p", 101),
			}, "password")
		})

		t.Run("update user email required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPut, ownerIDPath, env.Data.OwnerAccess, map[string]any{
				"username": env.Data.OwnerUsername,
			}, "email")
		})
		t.Run("update user email format", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPut, ownerIDPath, env.Data.OwnerAccess, map[string]any{
				"email":    "invalid-email",
				"username": env.Data.OwnerUsername,
			}, "email")
		})
		t.Run("update user email max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPut, ownerIDPath, env.Data.OwnerAccess, map[string]any{
				"email":    longEmail(96),
				"username": env.Data.OwnerUsername,
			}, "email")
		})
		t.Run("update user username min length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPut, ownerIDPath, env.Data.OwnerAccess, map[string]any{
				"email":    env.Data.OwnerEmail,
				"username": "",
			}, "username")
		})
		t.Run("update user username max length", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPut, ownerIDPath, env.Data.OwnerAccess, map[string]any{
				"email":    env.Data.OwnerEmail,
				"username": strings.Repeat("u", 101),
			}, "username")
		})

		t.Run("update user role enum required", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPatch, ownerIDPath+"/role", env.Data.AdminAccess, map[string]any{
				"role": "USER_ROLE_UNKNOWN",
			}, "role")
		})

		t.Run("get users limit lte 100", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodGet, usersPath+"?page=1&limit=101", env.Data.AdminAccess, nil, "limit")
		})

		t.Run("refresh token has no grpc validation rules", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, refreshPath, "", map[string]any{"refresh_token": "invalid"}, nil)
			if status == http.StatusBadRequest {
				t.Fatalf("unexpected grpc validation on refresh token endpoint, body=%s", string(body))
			}
		})
	})

	t.Run("auth grpc domain errors and statuses", func(t *testing.T) {
		usersPath := "/api/v1/users"
		ownerIDPath := usersPath + "/" + strconv.FormatInt(env.Data.OwnerID, 10)

		t.Run("401 missing authorization header", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, usersPath+"?page=1&limit=10", "", nil, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "authorization header is required")
		})

		t.Run("401 invalid jwt token", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, usersPath+"?page=1&limit=10", "invalid.token.value", nil, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "invalid token")
		})

		t.Run("401 invalid credentials from auth grpc login", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, "/api/v1/auth/login", "", map[string]any{
				"email":    env.Data.OwnerEmail,
				"password": env.Data.Password + "_wrong",
			}, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "invalid credentials")
		})

		t.Run("401 invalid refresh token from auth grpc", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, "/api/v1/auth/refresh", "", map[string]any{
				"refresh_token": "invalid-refresh-token",
			}, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "invalid token")
		})

		t.Run("409 register duplicate email", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, "/api/v1/auth/register", "", map[string]any{
				"email":    env.Data.OwnerEmail,
				"password": env.Data.Password,
			}, nil)
			env.RequireError(status, http.StatusConflict, body, "username or email already exists")
		})

		t.Run("409 create user duplicate email", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email":    env.Data.OwnerEmail,
				"username": env.Data.ManagedUsername + "-dup",
				"password": env.Data.Password,
			}, nil)
			env.RequireError(status, http.StatusConflict, body, "username or email already exists")
		})

		t.Run("404 get unknown user", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, usersPath+"/999999999", env.Data.AdminAccess, nil, nil)
			env.RequireError(status, http.StatusNotFound, body, "user not found")
		})

		t.Run("404 delete unknown user", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodDelete, usersPath+"/999999999", env.Data.AdminAccess, nil, nil)
			env.RequireError(status, http.StatusNotFound, body, "user not found")
		})

		t.Run("400 password hashing error on register (>72 bytes)", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, "/api/v1/auth/register", "", map[string]any{
				"email":    "hashing-register@example.com",
				"password": strings.Repeat("x", 73),
			}, nil)
			env.RequireError(status, http.StatusBadRequest, body, "error hashing password")
		})

		t.Run("400 password hashing error on create user (>72 bytes)", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, usersPath, env.Data.AdminAccess, map[string]any{
				"email":    "hashing-create-user@example.com",
				"username": "hashing-create-user",
				"password": strings.Repeat("x", 73),
			}, nil)
			env.RequireError(status, http.StatusBadRequest, body, "error hashing password")
		})

		t.Run("403 owner cannot create user", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, usersPath, env.Data.OwnerAccess, map[string]any{
				"email":    "owner-forbidden-create@example.com",
				"username": "owner-forbidden-create",
				"password": env.Data.Password,
			}, nil)
			env.RequireError(status, http.StatusForbidden, body, "forbidden")
		})

		t.Run("403 owner cannot update role", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPatch, ownerIDPath+"/role", env.Data.OwnerAccess, map[string]any{
				"role": "USER_ROLE_ADMIN",
			}, nil)
			env.RequireError(status, http.StatusForbidden, body, "forbidden")
		})

		t.Run("inactive user cannot get tokens", func(t *testing.T) {
			t.Run("deactivate owner", func(t *testing.T) {
				var resp map[string]any
				status, body := env.RequestJSON(
					http.MethodPatch,
					ownerIDPath+"/activity",
					env.Data.AdminAccess,
					map[string]any{"is_active": false},
					&resp,
				)
				env.RequireStatus(status, http.StatusOK, body)
			})

			t.Run("403 inactive user login forbidden", func(t *testing.T) {
				status, body := env.RequestJSON(http.MethodPost, "/api/v1/auth/login", "", map[string]any{
					"email":    env.Data.OwnerEmail,
					"password": env.Data.Password,
				}, nil)
				env.RequireError(status, http.StatusForbidden, body, "forbidden")
			})

			t.Run("403 inactive user refresh forbidden", func(t *testing.T) {
				status, body := env.RequestJSON(http.MethodPost, "/api/v1/auth/refresh", "", map[string]any{
					"refresh_token": env.Data.OwnerRefresh,
				}, nil)
				env.RequireError(status, http.StatusForbidden, body, "forbidden")
			})

			t.Run("reactivate owner", func(t *testing.T) {
				var resp map[string]any
				status, body := env.RequestJSON(
					http.MethodPatch,
					ownerIDPath+"/activity",
					env.Data.AdminAccess,
					map[string]any{"is_active": true},
					&resp,
				)
				env.RequireStatus(status, http.StatusOK, body)
			})
		})
	})
}

func assertValidationFields(
	t *testing.T,
	env *fixtures.Env,
	method, path, bearerToken string,
	body any,
	expectedFields ...string,
) {
	t.Helper()

	status, respBody := env.RequestJSON(method, path, bearerToken, body, nil)
	env.RequireValidationError(status, respBody, expectedFields...)
}

func longEmail(localLen int) string {
	if localLen < 1 {
		localLen = 1
	}
	return strings.Repeat("a", localLen) + "@e.com"
}
