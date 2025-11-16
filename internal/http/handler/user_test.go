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
	"auth_service/package/utils/errs"
)

func TestUserCreate(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful user creation",
			requestBody:    userReq,
			mockSetup:      mockUserCreateSuccess,
			expectedStatus: http.StatusCreated,
			respCheckers:   checkSuccessUserCreateResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkMessageError(errs.InvalidJSON),
		},
		{
			name:           "Email and Password required",
			requestBody:    request.UserCreate{},
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkFieldsRequired("Email", "Password"),
		},
		{
			name:           "Invalid Email and Password",
			requestBody:    loginBadEmailAndPasswordReq,
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			respCheckers: checkFieldsInvalid(map[string]error{
				"Email":    errs.InvalidEmail,
				"Password": errs.InvalidPassword,
			}),
		},
		{
			name:           "Email or username already exists",
			requestBody:    userReq,
			mockSetup:      mockUserCreateConflict,
			expectedStatus: http.StatusConflict,
			respCheckers:   checkMessageError(errs.UniqueUserField),
		},
		{
			name:           "Internal server error",
			requestBody:    userReq,
			mockSetup:      mockUserCreateServerError,
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   checkMessageError(errs.InternalServer),
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
			c.respCheckers(t, w)
			mockSvc.AssertExpectations(t)
		})
	}
}

func TestUserList(t *testing.T) {
	cases := []struct {
		name           string
		mockSetup      func(*mocks.Service)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful retrieving users",
			mockSetup:      mockUserListSuccess,
			expectedStatus: http.StatusOK,
			respCheckers:   checkSuccessUserListResponse(),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewService(t)
			c.mockSetup(mockSvc)

			var body []byte
			req := httptest.NewRequest("GET", "/users", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{svc: mockSvc}
			handler.UserList(w, req)

			mockSvc.AssertExpectations(t)
			assert.Equal(t, c.expectedStatus, w.Code)
			c.respCheckers(t, w)
		})
	}
}
