package mapper

import (
	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

func UserResponseFromProto(user *userv1.User) *dto.UserResponse {
	if user == nil {
		return nil
	}

	resp := &dto.UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role.String(),
		IsActive: user.IsActive,
	}

	if user.Username != nil {
		resp.Username = *user.Username
	}
	if user.CreatedAt != nil {
		resp.CreatedAt = user.CreatedAt.AsTime()
	}
	if user.UpdatedAt != nil {
		resp.UpdatedAt = user.UpdatedAt.AsTime()
	}

	return resp
}

func UsersResponseFromProto(resp *userv1.GetUsersResponse) *dto.UsersResponse {
	if resp == nil {
		return nil
	}

	users := make([]*dto.UserResponse, len(resp.Users))
	for i, user := range resp.Users {
		users[i] = userShortResponseFromProto(user)
	}

	return &dto.UsersResponse{Users: users}
}

func UpdateUserResponseFromProto(user *userv1.UpdateUser) *dto.UserResponse {
	if user == nil {
		return nil
	}

	resp := &dto.UserResponse{Email: user.Email}
	if user.Username != "" {
		resp.Username = user.Username
	}

	return resp
}

func userShortResponseFromProto(user *userv1.UserShort) *dto.UserResponse {
	if user == nil {
		return nil
	}

	resp := &dto.UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role.String(),
		IsActive: user.IsActive,
	}
	if user.Username != nil {
		resp.Username = *user.Username
	}

	return resp
}
