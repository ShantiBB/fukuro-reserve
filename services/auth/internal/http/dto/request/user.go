package request

type UserCreate struct {
	Username *string `json:"username"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
}

type UserUpdate struct {
	Username *string `json:"username"`
	Email    string  `json:"email" validate:"email"`
}

type UserUpdateRole struct {
	Role string `json:"role" validate:"required"`
}

type UserUpdateStatus struct {
	IsActive bool `json:"is_active" validate:"required"`
}
