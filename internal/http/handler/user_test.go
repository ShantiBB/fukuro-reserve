package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth_service/internal/http/lib/schemas/request"
	"auth_service/internal/mocks"
)

func TestUserCreate(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful user creation",
			requestBody:    userReq,
			mockSetup:      mockUserCreateSuccess,
			expectedStatus: http.StatusCreated,
			checkResponse:  checkSuccessResponse,
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			checkResponse:  checkInvalidJSONResponse,
		},
		{
			name:           "Email and Password required",
			requestBody:    request.UserCreate{},
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			checkResponse:  checkEmailAndPasswordRequired,
		},
		{
			name:           "Invalid Email and Password",
			requestBody:    loginBadEmailAndPasswordReq,
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			checkResponse:  checkInvalidEmailAndPassword,
		},
		{
			name:           "Email or username already exists",
			requestBody:    userReq,
			mockSetup:      mockUserCreateConflict,
			expectedStatus: http.StatusConflict,
			checkResponse:  checkConflictResponse,
		},
		{
			name:           "Internal server error",
			requestBody:    userReq,
			mockSetup:      mockUserCreateServerError,
			expectedStatus: http.StatusInternalServerError,
			checkResponse:  checkServerErrorResponse,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewService(t)
			c.mockSetup(mockSvc)

			var body []byte
			if str, ok := c.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(c.requestBody)
			}

			req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{svc: mockSvc}
			handler.UserCreate(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.checkResponse(t, w)

			mockSvc.AssertExpectations(t)
		})
	}
}
