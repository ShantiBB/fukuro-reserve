package request

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginByEmail struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
