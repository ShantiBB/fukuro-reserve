package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth/internal/http/dto/request"
	"auth/internal/mocks"
	"auth/pkg/utils/consts"
	"auth/pkg/utils/jwt"
	"auth/test/handler/unit"
)

func TestRegisterByEmail(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    any
		mockSetup      func(*mocks.MockService)
		expectedStatus int
		respCheckers   unit.ResponseChecker
	}{
		{
			name:           "Successful registration",
			requestBody:    unit.RegisterReq,
			mockSetup:      unit.MockRegisterSuccess,
			expectedStatus: http.StatusCreated,
			respCheckers:   unit.CheckSuccessTokenResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "",
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckMessageError(consts.InvalidJSON),
		},
		{
			name:           "Email and Password required",
			requestBody:    request.UserCreate{},
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckFieldsRequired("email", "password"),
		},
		{
			name:           "Invalid Email and Password",
			requestBody:    unit.LoginBadEmailAndPasswordReq,
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers: unit.CheckFieldsInvalid(map[string]error{
				"email":    consts.InvalidEmail,
				"password": consts.InvalidPassword,
			}),
		},
		{
			name:           "Email already exists",
			requestBody:    unit.RegisterReq,
			mockSetup:      unit.MockRegisterConflict,
			expectedStatus: http.StatusConflict,
			respCheckers:   unit.CheckMessageError(consts.UniqueUserField),
		},
		{
			name:           "Internal server error during registration",
			requestBody:    unit.RegisterReq,
			mockSetup:      unit.MockRegisterServerError,
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   unit.CheckMessageError(consts.InternalServer),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewMockService(t)
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
		requestBody    any
		mockSetup      func(*mocks.MockService)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful login",
			requestBody:    unit.LoginReq,
			mockSetup:      unit.MockLoginSuccess,
			expectedStatus: http.StatusOK,
			respCheckers:   unit.CheckSuccessTokenResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "",
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckMessageError(consts.InvalidJSON),
		},
		{
			name:           "Email and Password required",
			requestBody:    request.UserCreate{},
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckFieldsRequired("email", "password"),
		},
		{
			name:           "Invalid Email",
			requestBody:    unit.LoginBadEmailAndPasswordReq,
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckFieldsInvalid(map[string]error{"email": consts.InvalidEmail}),
		},
		{
			name:           "Invalid credentials",
			requestBody:    unit.LoginReq,
			mockSetup:      unit.MockLoginInvalidCredentials,
			expectedStatus: http.StatusUnauthorized,
			respCheckers:   unit.CheckMessageError(consts.InvalidCredentials),
		},
		{
			name:           "User not found",
			requestBody:    unit.LoginReq,
			mockSetup:      unit.MockLoginUserNotFound,
			expectedStatus: http.StatusUnauthorized,
			respCheckers:   unit.CheckMessageError(consts.InvalidCredentials),
		},
		{
			name:           "Internal server error",
			requestBody:    unit.LoginReq,
			mockSetup:      unit.MockLoginServerError,
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   unit.CheckMessageError(consts.InternalServer),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewMockService(t)
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
		requestBody    any
		mockSetup      func(*mocks.MockService)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful tokenCreds refresh",
			requestBody:    unit.RefreshReq,
			mockSetup:      unit.MockRefreshSuccess,
			expectedStatus: http.StatusOK,
			respCheckers:   unit.CheckSuccessTokenResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "",
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckMessageError(consts.InvalidJSON),
		},
		{
			name:           "Refresh tokenCreds required",
			requestBody:    jwt.RefreshToken{},
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckFieldsRequired("refresh_token"),
		},
		{
			name:           "Invalid refresh tokenCreds",
			requestBody:    unit.RefreshReq,
			mockSetup:      unit.MockRefreshInvalidToken,
			expectedStatus: http.StatusUnauthorized,
			respCheckers:   unit.CheckMessageError(consts.InvalidRefreshToken),
		},
		{
			name:           "Internal server error during tokenCreds refresh",
			requestBody:    unit.RefreshReq,
			mockSetup:      unit.MockRefreshServerError,
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   unit.CheckMessageError(consts.InternalServer),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewMockService(t)
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
