package dto

import userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"

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

func TokenResponseFromProto(resp *userv1.RegisterUserResponse) *TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}

	return &TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}

func TokenResponseFromLoginProto(resp *userv1.LoginUserResponse) *TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}

	return &TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}

func TokenResponseFromRefreshProto(resp *userv1.RefreshTokenResponse) *TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}

	return &TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}
