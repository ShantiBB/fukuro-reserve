package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth/internal/http/dto/response"
	"fukuro-reserve/pkg/utils/consts"
)

type ResponseChecker func(*testing.T, *httptest.ResponseRecorder)

var (
	checkSuccessUserResponse = func() ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.User
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, resp.ID, userMock.ID)
			assert.Equal(t, resp.Email, userMock.Email)
			assert.Equal(t, resp.Username, userMock.Username)
			assert.Equal(t, resp.Role, userMock.Role)
			assert.Equal(t, resp.IsActive, userMock.IsActive)
			assert.NotEmpty(t, resp.CreatedAt)
			assert.NotEmpty(t, resp.UpdatedAt)
		}
	}

	checkSuccessUserGetAllResponse = func() ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp []response.User
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, resp[0].ID, userMock.ID)
			assert.Equal(t, resp[0].Email, userMock.Email)
			assert.Equal(t, resp[0].Username, userMock.Username)
			assert.Equal(t, resp[0].Role, userMock.Role)
			assert.Equal(t, resp[0].IsActive, userMock.IsActive)
			assert.NotEmpty(t, resp[0].CreatedAt)
			assert.NotEmpty(t, resp[0].UpdatedAt)
		}
	}

	checkSuccessTokenResponse = func() ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.Token
			assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
			assert.Equal(t, tokensMock.Access, resp.Access)
			assert.Equal(t, tokensMock.Refresh, resp.Refresh)
		}
	}

	checkMessageError = func(expectedErr error) ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.ErrorSchema
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, expectedErr.Error(), resp.Message)
		}
	}

	checkFieldsRequired = func(fields ...string) ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.ValidateError
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			for _, field := range fields {
				assert.Equal(t, consts.FieldRequired.Error(), resp.Errors[field])
			}
		}
	}

	checkFieldsInvalid = func(fields map[string]error) ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.ValidateError
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			for field, expectedErr := range fields {
				assert.Equal(t, expectedErr.Error(), resp.Errors[field])
			}
		}
	}
)
