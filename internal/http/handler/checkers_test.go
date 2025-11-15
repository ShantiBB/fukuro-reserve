package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth_service/package/utils/errs"
)

var (
	checkSuccessResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, float64(1), response["id"])
		assert.Equal(t, "test@example.com", response["email"])
		assert.Equal(t, "test-user", response["username"])
	}

	checkTokenResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "access-token", resp["access_token"])
		assert.Equal(t, "refresh-token", resp["refresh_token"])
		assert.Equal(t, "Bearer", resp["token_type"])
	}

	checkEmailAndPasswordRequired = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["Email"])
		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["Password"])
	}

	checkInvalidEmailAndPassword = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.InvalidEmail.Error(), resp.Errors["Email"])
		assert.Equal(t, errs.InvalidPassword.Error(), resp.Errors["Password"])
	}

	checkLoginInvalidEmail = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.InvalidEmail.Error(), resp.Errors["Email"])
	}

	checkRefreshTokenRequired = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["RefreshToken"])
	}

	checkInvalidJSONResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["message"], errs.InvalidJSON.Error())
	}

	checkConflictResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["message"], errs.UniqueUserField.Error())
	}

	checkUnauthorizedResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["message"])
	}

	checkServerErrorResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["message"], errs.InternalServer.Error())
	}
)
