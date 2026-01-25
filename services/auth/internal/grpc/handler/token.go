package handler

import (
	"context"
	"log/slog"

	userv1 "auth/api/user/v1"
	"auth/internal/grpc/lib/utils/helper"
	"auth/internal/grpc/lib/utils/mapper"
)

func (h *Handler) RegisterUser(
	ctx context.Context,
	req *userv1.RegisterUserRequest,
) (*userv1.RegisterUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	user, err := mapper.RegisterUserRequestToDomain(req)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	created, err := h.svc.RegisterByEmail(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.RegisterUserResponse{
		Tokens: mapper.JWTTokenResponseToProto(created),
	}, nil
}

func (h *Handler) LoginUser(
	ctx context.Context,
	req *userv1.LoginUserRequest,
) (*userv1.LoginUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	user := mapper.LoginUserRequestToDomain(req)
	login, err := h.svc.LoginByEmail(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.LoginUserResponse{
		Tokens: mapper.JWTTokenResponseToProto(login),
	}, nil
}

func (h *Handler) RefreshToken(
	ctx context.Context,
	req *userv1.RefreshTokenRequest,
) (*userv1.RefreshTokenResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	refresh := mapper.RefreshTokenRequestToDomain(req)
	tokens, err := h.svc.RefreshToken(refresh)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.RefreshTokenResponse{
		Tokens: mapper.JWTTokenResponseToProto(tokens),
	}, nil
}
