package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	request2 "auth/internal/http/lib/schemas/request"
	"auth/internal/mocks"
	"fukuro-reserve/pkg/utils/errs"
	"fukuro-reserve/pkg/utils/jwt"
)

func TestRegisterByEmail(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		respCheckers   ResponseChecker
	}{
		{
			name:           "Successful registration",
			requestBody:    registerReq,
			mockSetup:      mockRegisterSuccess,
			expectedStatus: http.StatusCreated,
			respCheckers:   checkSuccessTokenResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "",
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkMessageError(errs.InvalidJSON),
		},
		{
			name:           "Email and Password required",
			requestBody:    request2.UserCreate{},
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkFieldsRequired("Email", "Password"),
		},
		{
			name:           "Invalid Email and Password",
			requestBody:    loginBadEmailAndPasswordReq,
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers: checkFieldsInvalid(map[string]error{
				"Email":    errs.InvalidEmail,
				"Password": errs.InvalidPassword,
			}),
		},
		{
			name:           "Email already exists",
			requestBody:    registerReq,
			mockSetup:      mockRegisterConflict,
			expectedStatus: http.StatusConflict,
			respCheckers:   checkMessageError(errs.UniqueEmailField),
		},
		{
			name:           "Internal server error during registration",
			requestBody:    registerReq,
			mockSetup:      mockRegisterServerError,
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

			req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{svc: mockSvc}
			handler.RegisterByEmail(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.respCheckers(t, w)

			mockSvc.AssertExpectations(t)
		})
	}
}

func TestLoginByEmail(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful login",
			requestBody:    loginReq,
			mockSetup:      mockLoginSuccess,
			expectedStatus: http.StatusOK,
			respCheckers:   checkSuccessTokenResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "",
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkMessageError(errs.InvalidJSON),
		},
		{
			name:           "Email and Password required",
			requestBody:    request2.UserCreate{},
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkFieldsRequired("Email", "Password"),
		},
		{
			name:           "Invalid Email",
			requestBody:    loginBadEmailAndPasswordReq,
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkFieldsInvalid(map[string]error{"Email": errs.InvalidEmail}),
		},
		{
			name:           "Invalid credentials",
			requestBody:    loginReq,
			mockSetup:      mockLoginInvalidCredentials,
			expectedStatus: http.StatusUnauthorized,
			respCheckers:   checkMessageError(errs.InvalidCredentials),
		},
		{
			name:           "User not found",
			requestBody:    loginReq,
			mockSetup:      mockLoginUserNotFound,
			expectedStatus: http.StatusUnauthorized,
			respCheckers:   checkMessageError(errs.InvalidCredentials),
		},
		{
			name:           "Internal server error",
			requestBody:    loginReq,
			mockSetup:      mockLoginServerError,
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

			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{svc: mockSvc}
			handler.LoginByEmail(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.respCheckers(t, w)

			mockSvc.AssertExpectations(t)
		})
	}
}

func TestRefreshToken(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful tokenCreds refresh",
			requestBody:    refreshReq,
			mockSetup:      mockRefreshSuccess,
			expectedStatus: http.StatusOK,
			respCheckers:   checkSuccessTokenResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "",
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkMessageError(errs.InvalidJSON),
		},
		{
			name:           "Refresh tokenCreds required",
			requestBody:    jwt.RefreshToken{},
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkFieldsRequired("RefreshToken"),
		},
		{
			name:           "Invalid refresh tokenCreds",
			requestBody:    refreshReq,
			mockSetup:      mockRefreshInvalidToken,
			expectedStatus: http.StatusUnauthorized,
			respCheckers:   checkMessageError(errs.InvalidRefreshToken),
		},
		{
			name:           "Internal server error during tokenCreds refresh",
			requestBody:    refreshReq,
			mockSetup:      mockRefreshServerError,
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

			req := httptest.NewRequest("POST", "/refresh", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{svc: mockSvc}
			handler.RefreshToken(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.respCheckers(t, w)

			mockSvc.AssertExpectations(t)
		})
	}
}
