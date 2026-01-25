package unit

import (
	"github.com/stretchr/testify/mock"

	"auth/internal/mocks"
	"auth/internal/repository/models"
	"auth/pkg/lib/utils/consts"
)

var (
	MockUserCreateSuccess = func(m *mocks.MockService) {
		m.On("InsertUser", mock.Anything, mock.Anything).Return(&UserMock, nil)
	}

	MockUserCreateConflict = func(m *mocks.MockService) {
		m.On("InsertUser", mock.Anything, mock.Anything).Return(nil, consts.ErrUniqueUserField)
	}

	MockUserCreateServerError = func(m *mocks.MockService) {
		m.On("InsertUser", mock.Anything, mock.Anything).Return(nil, consts.ErrInternalServer)
	}
)

var (
	MockUserGetAllSuccess = func(m *mocks.MockService) {
		m.On(
			"SelectUsers",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			&models.UserList{
				Users:      []*models.UserShort{UserShortMock},
				TotalCount: 1,
			}, nil,
		)
	}

	MockUserGetAllServerError = func(m *mocks.MockService) {
		m.On("SelectUsers", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, consts.ErrInternalServer)
	}
)

var (
	MockUserGetByIDSuccess = func(m *mocks.MockService) {
		m.On("SelectUserByID", mock.Anything, mock.Anything).Return(&UserMock, nil)
	}

	MockUserGetByIDNotFound = func(m *mocks.MockService) {
		m.On("SelectUserByID", mock.Anything, mock.Anything).Return(nil, consts.ErrUserNotFound)
	}

	MockUserGetByIDServerError = func(m *mocks.MockService) {
		m.On("SelectUserByID", mock.Anything, mock.Anything).Return(nil, consts.ErrInternalServer)
	}
)
