package handler

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"auth/internal/repository/models"
	"fukuro-reserve/pkg/utils/errs"
)

var (
	mockUserCreateSuccess = func(m *mocks.Service) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(&userMock, nil)
	}

	mockUserCreateConflict = func(m *mocks.Service) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, errs.UniqueEmailField)
	}

	mockUserCreateServerError = func(m *mocks.Service) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, errs.InternalServer)
	}
)

var (
	mockUserListSuccess = func(m *mocks.Service) {
		m.On("UserList", mock.Anything).Return([]models.User{{
			ID:        userMock.ID,
			Email:     userMock.Email,
			Username:  userMock.Username,
			Role:      userMock.Role,
			IsActive:  userMock.IsActive,
			CreatedAt: userMock.CreatedAt,
			UpdatedAt: userMock.UpdatedAt,
		}}, nil)
	}

	mockUserListServerError = func(m *mocks.Service) {
		m.On("UserList", mock.Anything).Return(nil, errs.InternalServer)
	}
)

var (
	mockUserGetByIDSuccess = func(m *mocks.Service) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(&userMock, nil)
	}

	mockUserGetByIDNotFound = func(m *mocks.Service) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(nil, errs.UserNotFound)
	}

	mockUserGetByIDServerError = func(m *mocks.Service) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(nil, errs.InternalServer)
	}
)
