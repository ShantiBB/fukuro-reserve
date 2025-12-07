package unit

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"auth/internal/repository/postgres/models"
	"fukuro-reserve/pkg/utils/consts"
)

var (
	MockUserCreateSuccess = func(m *mocks.MockService) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(&UserMock, nil)
	}

	MockUserCreateConflict = func(m *mocks.MockService) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, consts.UniqueEmailField)
	}

	MockUserCreateServerError = func(m *mocks.MockService) {
		m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, consts.InternalServer)
	}
)

var (
	MockUserGetAllSuccess = func(m *mocks.MockService) {
		m.On("UserGetAll",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&models.UserList{
			Users:      []models.UserShort{UserShortMock},
			TotalCount: 1,
		}, nil)
	}

	MockUserGetAllServerError = func(m *mocks.MockService) {
		m.On("UserGetAll", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.InternalServer)
	}
)

var (
	MockUserGetByIDSuccess = func(m *mocks.MockService) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(&UserMock, nil)
	}

	MockUserGetByIDNotFound = func(m *mocks.MockService) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(nil, consts.UserNotFound)
	}

	MockUserGetByIDServerError = func(m *mocks.MockService) {
		m.On("UserGetByID", mock.Anything, mock.Anything).Return(nil, consts.InternalServer)
	}
)
