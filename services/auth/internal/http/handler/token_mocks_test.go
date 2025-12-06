package handler

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"fukuro-reserve/pkg/utils/consts"
)

// RegisterByEmail mocks
var (
	mockRegisterSuccess = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&tokensMock, nil)
	}

	mockRegisterConflict = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.UniqueEmailField)
	}

	mockRegisterServerError = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.InternalServer)
	}
)

// LoginByEmail mocks
var (
	mockLoginSuccess = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&tokensMock, nil)
	}

	mockLoginInvalidCredentials = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.InvalidCredentials)
	}

	mockLoginUserNotFound = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.UserNotFound)
	}

	mockLoginServerError = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.InternalServer)
	}
)

// RefreshToken mocks
var (
	mockRefreshSuccess = func(m *mocks.MockService) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(&tokensMock, nil)
	}

	mockRefreshInvalidToken = func(m *mocks.MockService) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, consts.InvalidRefreshToken)
	}

	mockRefreshServerError = func(m *mocks.MockService) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, consts.InternalServer)
	}
)
