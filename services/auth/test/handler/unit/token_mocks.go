package unit

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"auth/pkg/utils/consts"
)

// RegisterByEmail mocks
var (
	MockRegisterSuccess = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&TokensMock, nil)
	}

	MockRegisterConflict = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.UniqueUserField)
	}

	MockRegisterServerError = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.InternalServer)
	}
)

// LoginByEmail mocks
var (
	MockLoginSuccess = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&TokensMock, nil)
	}

	MockLoginInvalidCredentials = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.InvalidCredentials)
	}

	MockLoginUserNotFound = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.UserNotFound)
	}

	MockLoginServerError = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.InternalServer)
	}
)

// RefreshToken mocks
var (
	MockRefreshSuccess = func(m *mocks.MockService) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(&TokensMock, nil)
	}

	MockRefreshInvalidToken = func(m *mocks.MockService) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, consts.InvalidRefreshToken)
	}

	MockRefreshServerError = func(m *mocks.MockService) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, consts.InternalServer)
	}
)
