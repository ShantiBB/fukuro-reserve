package response

import (
	"time"

	"auth/internal/http/utils/pagination"
	"auth/internal/repository/models"
)

type User struct {
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Username  *string         `json:"username"`
	Email     string          `json:"email"`
	Role      models.UserRole `json:"role"`
	ID        int64           `json:"id"`
	IsActive  bool            `json:"is_active"`
}

type UserShort struct {
	Username *string         `json:"username"`
	Email    string          `json:"email"`
	Role     models.UserRole `json:"role"`
	ID       int64           `json:"id"`
	IsActive bool            `json:"is_active"`
}

type UserList struct {
	Links           pagination.Links `json:"links"`
	Users           []UserShort      `json:"users"`
	CurrentPage     uint64           `json:"current_page"`
	Limit           uint64           `json:"limit"`
	TotalPageCount  uint64           `json:"total_page_count"`
	TotalUsersCount uint64           `json:"total_users_count"`
}

type UserRoleStatus struct {
	Message string `json:"message"`
	Role    string `json:"role"`
}

type UserActiveStatus struct {
	Message  string `json:"message"`
	IsActive bool   `json:"is_active"`
}
