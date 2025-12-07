package request

type UserCreate struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"strongpass123!#"`
}

type UserUpdate struct {
	Username *string `json:"username" example:"username"`
	Email    string  `json:"email" validate:"email" example:"user@example.com"`
}

type UserRoleStatus struct {
	Role string `json:"role" validate:"required" example:"user"`
}

type UserActiveStatus struct {
	IsActive *bool `json:"is_active" validate:"required"`
}

type PaginationQuery struct {
	Page  uint64
	Limit uint64
}
