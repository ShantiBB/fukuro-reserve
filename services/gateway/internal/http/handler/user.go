package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	httpauth "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/auth"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/request"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
)

// GetUsers godoc
// @Summary       Get all users
// @Description   Returns users with pagination. Requires JWT auth.
// @Tags          users
// @Produce       json
// @Security      Bearer
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.UsersResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /auth/users [get]
func (h *AuthHandler) GetUsers(c *gin.Context) {
	ctx, err := httpauth.OutgoingContextWithAuthorization(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	page, err := request.OptionalUint64Query(c, "page")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrPageMustBePositiveInteger})
		return
	}
	limit, err := request.OptionalUint64Query(c, "limit")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrLimitMustBePositiveInteger})
		return
	}

	resp, err := h.service.GetUsers(ctx, page, limit)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateUser godoc
// @Summary       Create a new user
// @Description   Creates a user account (admin/staff operation). Requires JWT auth.
// @Tags          users
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         request body dto.CreateUserRequest true "Create user request"
// @Success       201 {object} dto.UserResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       409 {object} responder.ErrorResponse
// @Router        /auth/users [post]
func (h *AuthHandler) CreateUser(c *gin.Context) {
	ctx, err := httpauth.OutgoingContextWithAuthorization(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	var req dto.CreateUserRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.CreateUser(ctx, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetUser godoc
// @Summary       Get user by ID
// @Description   Returns a user by ID. Requires JWT auth.
// @Tags          users
// @Produce       json
// @Security      Bearer
// @Param         id path int true "User ID" minimum(1) example(1)
// @Success       200 {object} dto.UserResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /auth/users/{id} [get]
func (h *AuthHandler) GetUser(c *gin.Context) {
	ctx, err := httpauth.OutgoingContextWithAuthorization(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := request.PositiveInt64Path(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidUserID})
		return
	}

	resp, err := h.service.GetUser(ctx, id)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary       Update user
// @Description   Updates user profile fields. Requires JWT auth.
// @Tags          users
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         id path int true "User ID" minimum(1) example(1)
// @Param         request body dto.UpdateUserRequest true "Update user request"
// @Success       200 {object} dto.UserResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /auth/users/{id} [put]
func (h *AuthHandler) UpdateUser(c *gin.Context) {
	ctx, err := httpauth.OutgoingContextWithAuthorization(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := request.PositiveInt64Path(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidUserID})
		return
	}

	var req dto.UpdateUserRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateUser(ctx, id, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUserActivity godoc
// @Summary       Update user activity status
// @Description   Updates active/inactive flag for a user. Requires JWT auth.
// @Tags          users
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         id path int true "User ID" minimum(1) example(1)
// @Param         request body dto.UpdateUserActivityRequest true "Update activity request"
// @Success       200 {object} dto.UpdateUserActivityResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /auth/users/{id}/activity [patch]
func (h *AuthHandler) UpdateUserActivity(c *gin.Context) {
	ctx, err := httpauth.OutgoingContextWithAuthorization(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := request.PositiveInt64Path(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidUserID})
		return
	}

	var req dto.UpdateUserActivityRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateUserActivity(ctx, id, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUserRole godoc
// @Summary       Update user role
// @Description   Updates user role. Requires JWT auth.
// @Tags          users
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         id path int true "User ID" minimum(1) example(1)
// @Param         request body dto.UpdateUserRoleRequest true "Update role request"
// @Success       200 {object} dto.UpdateUserRoleResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /auth/users/{id}/role [patch]
func (h *AuthHandler) UpdateUserRole(c *gin.Context) {
	ctx, err := httpauth.OutgoingContextWithAuthorization(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := request.PositiveInt64Path(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidUserID})
		return
	}

	var req dto.UpdateUserRoleRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateUserRole(ctx, id, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary       Delete user
// @Description   Deletes user by ID. Requires JWT auth.
// @Tags          users
// @Security      Bearer
// @Param         id path int true "User ID" minimum(1) example(1)
// @Success       204
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /auth/users/{id} [delete]
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	ctx, err := httpauth.OutgoingContextWithAuthorization(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := request.PositiveInt64Path(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidUserID})
		return
	}

	if err = h.service.DeleteUser(ctx, id); err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
