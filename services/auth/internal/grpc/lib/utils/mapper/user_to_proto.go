package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	userv1 "auth/api/user/v1"
	"auth/internal/repository/models"
)

func userRoleToProto(status models.UserRole) userv1.UserRole {
	var s userv1.UserRole
	switch status {
	case models.UserRoleUser:
		s = userv1.UserRole_USER_ROLE_USER
	case models.UserRoleModerator:
		s = userv1.UserRole_USER_ROLE_MODERATOR
	case models.UserRoleAdmin:
		s = userv1.UserRole_USER_ROLE_ADMIN
	default:
		s = userv1.UserRole_USER_ROLE_UNSPECIFIED
	}
	return s
}

func CreateUserResponseToProto(resp *models.User) *userv1.User {
	return &userv1.User{
		Id:        resp.ID,
		Username:  resp.Username,
		Email:     resp.Email,
		Role:      userRoleToProto(resp.Role),
		IsActive:  resp.IsActive,
		CreatedAt: timestamppb.New(resp.CreatedAt),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}
}

func UserResponseToProto(resp *models.User) *userv1.User {
	return &userv1.User{
		Id:        resp.ID,
		Username:  resp.Username,
		Email:     resp.Email,
		Role:      userRoleToProto(resp.Role),
		IsActive:  resp.IsActive,
		CreatedAt: timestamppb.New(resp.CreatedAt),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}
}

func UpdateUserResponseToProto(resp *models.UpdateUser) *userv1.UpdateUser {
	return &userv1.UpdateUser{
		Username: resp.Username,
		Email:    resp.Email,
	}
}

func UserShortResponseToProto(resp *models.UserShort) *userv1.UserShort {
	return &userv1.UserShort{
		Id:       resp.ID,
		Username: resp.Username,
		Email:    resp.Email,
		Role:     userRoleToProto(resp.Role),
		IsActive: resp.IsActive,
	}
}

func UsersResponseToProto(resp []*models.UserShort) []*userv1.UserShort {
	users := make([]*userv1.UserShort, len(resp))
	for i, u := range resp {
		users[i] = UserShortResponseToProto(u)
	}

	return users
}
