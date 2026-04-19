package dto

type RegisterRequest struct {
	Email    string `json:"email" example:"guest@example.com"`
	Password string `json:"password" example:"Passw0rd!123"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"guest@example.com"`
	Password string `json:"password" example:"Passw0rd!123"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.refresh.token"`
}

type TokenResponse struct {
	Access  string `json:"access" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.access.token"`
	Refresh string `json:"refresh" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.refresh.token"`
}
