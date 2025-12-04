package handler

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"fukuro-reserve/pkg/utils/errs"
)

// RegisterByEmail mocks
var (
	mockRegisterSuccess = func(m *mocks.Service) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&tokensMock, nil)
	}

	mockRegisterConflict = func(m *mocks.Service) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errs.UniqueEmailField)
	}

	mockRegisterServerError = func(m *mocks.Service) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errs.InternalServer)
	}
)

// LoginByEmail mocks
var (
	mockLoginSuccess = func(m *mocks.Service) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&tokensMock, nil)
	}

	mockLoginInvalidCredentials = func(m *mocks.Service) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errs.InvalidCredentials)
	}

	mockLoginUserNotFound = func(m *mocks.Service) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errs.UserNotFound)
	}

	mockLoginServerError = func(m *mocks.Service) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errs.InternalServer)
	}
)

// RefreshToken mocks
var (
	mockRefreshSuccess = func(m *mocks.Service) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(&tokensMock, nil)
	}

	mockRefreshInvalidToken = func(m *mocks.Service) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, errs.InvalidRefreshToken)
	}

	mockRefreshServerError = func(m *mocks.Service) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, errs.InternalServer)
	}
)
