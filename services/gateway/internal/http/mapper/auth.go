package mapper

import (
	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

func TokenResponseFromRegister(resp *userv1.RegisterUserResponse) *dto.TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}

	return &dto.TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}

func TokenResponseFromLogin(resp *userv1.LoginUserResponse) *dto.TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}

	return &dto.TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}

func TokenResponseFromRefresh(resp *userv1.RefreshTokenResponse) *dto.TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}

	return &dto.TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}
