package dto

import "time"

// CreateUserRequest User DTOs.
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
