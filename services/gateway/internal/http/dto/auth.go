package dto

// Auth DTOs.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
