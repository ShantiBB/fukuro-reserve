package handler

import (
	"context"
	"log/slog"

	userv1 "auth/api/user/v1"
	helper2 "auth/internal/grpc/lib/utils/helper"
	mapper2 "auth/internal/grpc/lib/utils/mapper"
)

func (h *Handler) CreateUser(
	ctx context.Context,
	req *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper2.HandleValidationErr(err)
	}

	user, err := mapper2.CreateUserRequestToDomain(req)
	if err != nil {
		return nil, helper2.HandleDomainErr(err)
	}

	created, err := h.svc.CreateUser(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper2.HandleDomainErr(err)
	}

	return &userv1.CreateUserResponse{
		User: mapper2.CreateUserResponseToProto(created),
	}, nil
}

func (h *Handler) GetUsers(
	ctx context.Context,
	req *userv1.GetUsersRequest,
) (*userv1.GetUsersResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper2.HandleValidationErr(err)
	}

	userList, err := h.svc.GetUsers(ctx, req.Page, req.Limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper2.HandleDomainErr(err)
	}

	return &userv1.GetUsersResponse{
		Users:      mapper2.UsersResponseToProto(userList.Users),
		TotalCount: userList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}

func (h *Handler) GetUser(
	ctx context.Context,
	req *userv1.GetUserRequest,
) (*userv1.GetUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper2.HandleValidationErr(err)
	}

	user, err := h.svc.GetUserByID(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper2.HandleDomainErr(err)
	}

	return &userv1.GetUserResponse{
		User: mapper2.UserResponseToProto(user),
	}, nil
}

func (h *Handler) UpdateUser(
	ctx context.Context,
	req *userv1.UpdateUserRequest,
) (*userv1.UpdateUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper2.HandleValidationErr(err)
	}

	user := mapper2.UpdateUserRequestToDomain(req)
	if err := h.svc.UpdateUserByID(ctx, user); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper2.HandleDomainErr(err)
	}

	return &userv1.UpdateUserResponse{
		User: mapper2.UpdateUserResponseToProto(user),
	}, nil
}

func (h *Handler) DeleteUser(
	ctx context.Context,
	req *userv1.DeleteUserRequest,
) (*userv1.DeleteUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper2.HandleValidationErr(err)
	}

	if err := h.svc.DeleteUserByID(ctx, req.Id); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper2.HandleDomainErr(err)
	}

	return &userv1.DeleteUserResponse{
		Message: "success",
	}, nil
}
