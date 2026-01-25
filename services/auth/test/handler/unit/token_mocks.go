package unit

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"auth/pkg/lib/utils/consts"
)

// RegisterByEmail mocks
var (
	MockRegisterSuccess = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&TokensMock, nil)
	}

	MockRegisterConflict = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.ErrUniqueUserField)
	}

	MockRegisterServerError = func(m *mocks.MockService) {
		m.On("RegisterByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.ErrInternalServer)
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
			Return(nil, consts.ErrInvalidCredentials)
	}

	MockLoginUserNotFound = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.ErrUserNotFound)
	}

	MockLoginServerError = func(m *mocks.MockService) {
		m.On("LoginByEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.ErrInternalServer)
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
			Return(nil, consts.ErrInvalidToken)
	}

	MockRefreshServerError = func(m *mocks.MockService) {
		m.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, consts.ErrInternalServer)
	}
)
