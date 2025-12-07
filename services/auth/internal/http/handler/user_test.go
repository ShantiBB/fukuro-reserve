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

	"auth/internal/http/dto/request"
	"auth/internal/mocks"
	"auth/test/handler/unit"
	"fukuro-reserve/pkg/utils/consts"
)

func TestUserCreate(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    any
		mockSetup      func(*mocks.MockService)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful user creation",
			requestBody:    unit.UserReq,
			mockSetup:      unit.MockUserCreateSuccess,
			expectedStatus: http.StatusCreated,
			respCheckers:   unit.CheckSuccessUserResponse(),
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			mockSetup:      func(*mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckMessageError(consts.InvalidJSON),
		},
		{
			name:           "Email and Password required",
			requestBody:    request.UserCreate{},
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckFieldsRequired("Email", "Password"),
		},
		{
			name:           "Invalid Email and Password",
			requestBody:    unit.LoginBadEmailAndPasswordReq,
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
			respCheckers: unit.CheckFieldsInvalid(map[string]error{
				"Email":    consts.InvalidEmail,
				"Password": consts.InvalidPassword,
			}),
		},
		{
			name:           "Email or username already exists",
			requestBody:    unit.UserReq,
			mockSetup:      unit.MockUserCreateConflict,
			expectedStatus: http.StatusConflict,
			respCheckers:   unit.CheckMessageError(consts.UniqueUserField),
		},
		{
			name:           "Internal server error",
			requestBody:    unit.UserReq,
			mockSetup:      unit.MockUserCreateServerError,
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

func TestUserGetAll(t *testing.T) {
	cases := []struct {
		name           string
		mockSetup      func(*mocks.MockService)
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful retrieving users",
			mockSetup:      unit.MockUserGetAllSuccess,
			expectedStatus: http.StatusOK,
			respCheckers:   unit.CheckSuccessUserGetAllResponse(),
		},
		{
			name:           "Internal server error",
			mockSetup:      unit.MockUserGetAllServerError,
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   unit.CheckMessageError(consts.InternalServer),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewMockService(t)
			c.mockSetup(mockSvc)

			var body []byte
			req := httptest.NewRequest("GET", "/users", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{svc: mockSvc}
			handler.UserGetAll(w, req)

			mockSvc.AssertExpectations(t)
			assert.Equal(t, c.expectedStatus, w.Code)
			c.respCheckers(t, w)
		})
	}
}

func TestUserGetByID(t *testing.T) {
	cases := []struct {
		name           string
		mockSetup      func(*mocks.MockService)
		userID         string
		expectedStatus int
		respCheckers   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful retrieving user",
			mockSetup:      unit.MockUserGetByIDSuccess,
			userID:         strconv.FormatInt(unit.UserMock.ID, 10),
			expectedStatus: http.StatusOK,
			respCheckers:   unit.CheckSuccessUserResponse(),
		},
		{
			name:           "User not found",
			mockSetup:      unit.MockUserGetByIDNotFound,
			userID:         "999",
			expectedStatus: http.StatusNotFound,
			respCheckers:   unit.CheckMessageError(consts.UserNotFound),
		},
		{
			name:           "Invalid user ID",
			mockSetup:      func(m *mocks.MockService) {},
			userID:         "abc",
			expectedStatus: http.StatusBadRequest,
			respCheckers:   unit.CheckMessageError(consts.InvalidID),
		},
		{
			name:           "Internal server error",
			mockSetup:      unit.MockUserGetByIDServerError,
			userID:         strconv.FormatInt(unit.UserMock.ID, 10),
			expectedStatus: http.StatusInternalServerError,
			respCheckers:   unit.CheckMessageError(consts.InternalServer),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewMockService(t)
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
