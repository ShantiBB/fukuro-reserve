package response

import (
	"time"

	"fukuro-reserve/pkg/utils/helper"
)

type User struct {
	ID        int64     `json:"id"`
	Username  *string   `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserShort struct {
	ID       int64   `json:"id"`
	Username *string `json:"username"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	IsActive bool    `json:"is_active"`
}

type UserList struct {
	Users           []UserShort            `json:"users"`
	CurrentPage     uint64                 `json:"current_page"`
	Limit           uint64                 `json:"limit"`
	Links           helper.PaginationLinks `json:"links"`
	TotalPageCount  uint64                 `json:"total_page_count"`
	TotalUsersCount uint64                 `json:"total_users_count"`
}

type UserRoleStatus struct {
	Message string `json:"message"`
	Role    string `json:"role"`
}

type UserActiveStatus struct {
	Message  string `json:"message"`
	IsActive bool   `json:"is_active"`
}
