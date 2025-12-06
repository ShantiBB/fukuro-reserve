package response

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  *string   `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserList struct {
	Users           []User `json:"users"`
	CurrentPage     uint64 `json:"current_page"`
	Limit           uint64 `json:"limit"`
	HasPrevPage     bool   `json:"has_prev_page"`
	HasNextPage     bool   `json:"has_next_page"`
	TotalPageCount  uint64 `json:"total_page_count"`
	TotalUsersCount uint64 `json:"total_users_count"`
}

type UserShort struct {
	ID       int64   `json:"id"`
	Username *string `json:"username"`
	Email    string  `json:"email"`
}
