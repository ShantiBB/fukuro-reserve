package handler

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"auth/internal/repository/postgres/models"
	"fukuro-reserve/pkg/utils/errs"
)

var (
	mockUserCreateSuccess = func(m *mocks.MockService) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(&userMock, nil)
	}

	mockUserCreateConflict = func(m *mocks.MockService) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, errs.UniqueEmailField)
	}

	mockUserCreateServerError = func(m *mocks.MockService) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, errs.InternalServer)
	}
)

var (
	mockUserGetAllSuccess = func(m *mocks.MockService) {
		m.On("UserGetAll", mock.Anything, uint64(100), uint64(0)).
			Return([]models.User{userMock}, nil)
	}

	mockUserGetAllServerError = func(m *mocks.MockService) {
		m.On("UserGetAll", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errs.InternalServer)
	}
)

var (
	mockUserGetByIDSuccess = func(m *mocks.MockService) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(&userMock, nil)
	}

	mockUserGetByIDNotFound = func(m *mocks.MockService) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(nil, errs.UserNotFound)
	}

	mockUserGetByIDServerError = func(m *mocks.MockService) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(nil, errs.InternalServer)
	}
)
