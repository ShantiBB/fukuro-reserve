package handler

import (
	"context"
	"log/slog"

	userv1 "auth/api/user/v1"
	"auth/internal/grpc/lib/utils/helper"
	"auth/internal/grpc/lib/utils/mapper"
)

func (h *Handler) CreateUser(
	ctx context.Context,
	req *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	user, err := mapper.CreateUserRequestToDomain(req)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	created, err := h.svc.CreateUser(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.CreateUserResponse{
		User: mapper.CreateUserResponseToProto(created),
	}, nil
}

func (h *Handler) GetUsers(
	ctx context.Context,
	req *userv1.GetUsersRequest,
) (*userv1.GetUsersResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	userList, err := h.svc.GetUsers(ctx, req.Page, req.Limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.GetUsersResponse{
		Users:      mapper.UsersResponseToProto(userList.Users),
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
		return nil, helper.HandleValidationErr(err)
	}

	user, err := h.svc.GetUserByID(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.GetUserResponse{
		User: mapper.UserResponseToProto(user),
	}, nil
}

func (h *Handler) UpdateUser(
	ctx context.Context,
	req *userv1.UpdateUserRequest,
) (*userv1.UpdateUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	user := mapper.UpdateUserRequestToDomain(req)
	if err := h.svc.UpdateUserByID(ctx, user); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.UpdateUserResponse{
		User: mapper.UpdateUserResponseToProto(user),
	}, nil
}

func (h *Handler) UpdateUserActivity(
	ctx context.Context,
	req *userv1.UpdateUserActivityRequest,
) (*userv1.UpdateUserActivityResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	isActive := req.IsActive.GetValue()
	if err := h.svc.UpdateUserActiveStatus(ctx, req.Id, isActive); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.UpdateUserActivityResponse{
		IsActive: isActive,
	}, nil
}

func (h *Handler) UpdateUserRole(
	ctx context.Context,
	req *userv1.UpdateUserRoleRequest,
) (*userv1.UpdateUserRoleResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	newRole := mapper.UpdateUserRoleRequestToDomain(req)
	if err := h.svc.UpdateUserRoleStatus(ctx, req.Id, newRole); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.UpdateUserRoleResponse{
		Role: req.Role,
	}, nil
}

func (h *Handler) DeleteUser(
	ctx context.Context,
	req *userv1.DeleteUserRequest,
) (*userv1.DeleteUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	if err := h.svc.DeleteUserByID(ctx, req.Id); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.DeleteUserResponse{
		Message: "success",
	}, nil
}
