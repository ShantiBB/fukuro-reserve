package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth_service/package/utils/errs"
)

var checkSuccessResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, float64(1), response["id"])
	assert.Equal(t, "test@example.com", response["email"])
	assert.Equal(t, "test-user", response["username"])
}

var checkInvalidJSONResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["message"], errs.InvalidJSON.Error())
}

var checkConflictResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["message"], errs.UniqueUserField.Error())
}

var checkServerErrorResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["message"], errs.InternalServer.Error())
}
