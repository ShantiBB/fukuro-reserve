package dto

import (
	"time"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
)

// User DTOs.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserActivityRequest struct {
	IsActive bool `json:"is_active"`
}

type UpdateUserActivityResponse struct {
	IsActive bool `json:"is_active"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role"`
}

type UpdateUserRoleResponse struct {
	Role string `json:"role"`
}

type UserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Username  string    `json:"username,omitempty"`
	Role      string    `json:"role"`
	Id        int64     `json:"id"`
	IsActive  bool      `json:"is_active"`
}

type UsersResponse struct {
	Users []*UserResponse `json:"users"`
}

func UserResponseFromProto(user *userv1.User) *UserResponse {
	if user == nil {
		return nil
	}

	resp := &UserResponse{
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

func userShortResponseFromProto(user *userv1.UserShort) *UserResponse {
	if user == nil {
		return nil
	}

	resp := &UserResponse{
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

func UsersResponseFromProto(resp *userv1.GetUsersResponse) *UsersResponse {
	if resp == nil {
		return nil
	}

	users := make([]*UserResponse, len(resp.Users))
	for i, u := range resp.Users {
		users[i] = userShortResponseFromProto(u)
	}

	return &UsersResponse{Users: users}
}

func UpdateUserResponseFromProto(user *userv1.UpdateUser) *UserResponse {
	if user == nil {
		return nil
	}

	resp := &UserResponse{Email: user.Email}
	if user.Username != "" {
		resp.Username = user.Username
	}

	return resp
}
