package dto

import "time"

type CreateUserRequest struct {
	Email    string `json:"email" example:"manager@example.com"`
	Username string `json:"username,omitempty" example:"manager01"`
	Password string `json:"password" example:"Passw0rd!123"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" example:"manager.updated@example.com"`
	Username string `json:"username" example:"manager02"`
}

type UpdateUserActivityRequest struct {
	IsActive bool `json:"is_active" example:"true"`
}

type UpdateUserActivityResponse struct {
	IsActive bool `json:"is_active" example:"true"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role" example:"ROLE_ADMIN"`
}

type UpdateUserRoleResponse struct {
	Role string `json:"role" example:"ROLE_ADMIN"`
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
