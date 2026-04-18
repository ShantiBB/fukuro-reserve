package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	httpauth "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/auth"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/request"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation"
)

// Register godoc
// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register request"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} responder.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateRegisterRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Login godoc
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} responder.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateLoginRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} responder.ErrorResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateRefreshTokenRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.RefreshToken(c.Request.Context(), req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} dto.UsersResponse
// @Router /auth/users [get]
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
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "Create user request"
// @Success 201 {object} dto.UserResponse
// @Router /auth/users [post]
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
	if err = validation.ValidateCreateUserRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
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
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Router /auth/users/{id} [get]
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
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRequest true "Update user request"
// @Success 200 {object} dto.UserResponse
// @Router /auth/users/{id} [put]
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
	if err = validation.ValidateUpdateUserRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
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
// @Summary Update user activity status
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserActivityRequest true "Update activity request"
// @Success 200 {object} dto.UpdateUserActivityResponse
// @Router /auth/users/{id}/activity [patch]
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
// @Summary Update user role
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRoleRequest true "Update role request"
// @Success 200 {object} dto.UpdateUserRoleResponse
// @Router /auth/users/{id}/role [patch]
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
	if err = validation.ValidateUpdateUserRoleRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
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
// @Summary Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 204
// @Router /auth/users/{id} [delete]
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
