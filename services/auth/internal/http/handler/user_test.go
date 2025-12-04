package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"auth/internal/http/lib/schemas/request"
	"auth/internal/mocks"
	"fukuro-reserve/pkg/utils/errs"
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
			respCheckers:   checkSuccessUserResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			mockSetup:      func(m *mocks.Service) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkMessageError(errs.InvalidJSON),
		},
		{
			name:           "Email and Password required",
			requestBody:    request.UserCreate{},
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
		{
			name:           "Internal server error",
			mockSetup:      mockUserListServerError,
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   checkMessageError(errs.InternalServer),
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

func TestUserGetByID(t *testing.T) {
	cases := []struct {
		name           string
		mockSetup      func(*mocks.Service)
		userID         string
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful retrieving user",
			mockSetup:      mockUserGetByIDSuccess,
			userID:         strconv.FormatInt(userMock.ID, 10),
			expectedStatus: http.StatusOK,
			respCheckers:   checkSuccessUserResponse(),
		},
		{
			name:           "User not found",
			mockSetup:      mockUserGetByIDNotFound,
			userID:         "999",
			expectedStatus: http.StatusNotFound,
			respCheckers:   checkMessageError(errs.UserNotFound),
		},
		{
			name:           "Invalid user ID",
			mockSetup:      func(m *mocks.Service) {},
			userID:         "abc",
			expectedStatus: http.StatusBadRequest,
			respCheckers:   checkMessageError(errs.InvalidID),
		},
		{
			name:           "Internal server error",
			mockSetup:      mockUserGetByIDServerError,
			userID:         strconv.FormatInt(userMock.ID, 10),
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   checkMessageError(errs.InternalServer),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewService(t)
			c.mockSetup(mockSvc)

			handler := &Handler{svc: mockSvc}

			req := httptest.NewRequest(http.MethodGet, "/users/"+c.userID, nil)
			w := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", c.userID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			handler.UserGetByID(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.respCheckers(t, w)
		})
	}
}
