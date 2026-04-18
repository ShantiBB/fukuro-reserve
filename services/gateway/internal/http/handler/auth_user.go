package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/query"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
)

// GetUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} dto.UsersResponse
// @Router /api/v1/auth/users [get]
func (h *AuthHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		responder.Error(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	page := query.ParseUint64(r.URL.Query().Get("page"))
	limit := query.ParseUint64(r.URL.Query().Get("limit"))

	resp, err := h.service.GetUsers(ctx, page, limit)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// CreateUser godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "Create user request"
// @Success 201 {object} dto.UserResponse
// @Router /api/v1/auth/users [post]
func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		responder.Error(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.CreateUser(ctx, req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusCreated, resp)
}

// GetUser godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Router /api/v1/auth/users/{id} [get]
func (h *AuthHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		responder.Error(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := query.ParseInt64(idStr)
	if id == 0 {
		responder.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	resp, err := h.service.GetUser(ctx, id)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRequest true "Update user request"
// @Success 200 {object} dto.UserResponse
// @Router /api/v1/auth/users/{id} [put]
func (h *AuthHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		responder.Error(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := query.ParseInt64(idStr)
	if id == 0 {
		responder.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.UpdateUser(ctx, id, req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// UpdateUserActivity godoc
// @Summary Update user activity status
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserActivityRequest true "Update activity request"
// @Success 200 {object} dto.UpdateUserActivityResponse
// @Router /api/v1/auth/users/{id}/activity [patch]
func (h *AuthHandler) UpdateUserActivity(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		responder.Error(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := query.ParseInt64(idStr)
	if id == 0 {
		responder.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req dto.UpdateUserActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.UpdateUserActivity(ctx, id, req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// UpdateUserRole godoc
// @Summary Update user role
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRoleRequest true "Update role request"
// @Success 200 {object} dto.UpdateUserRoleResponse
// @Router /api/v1/auth/users/{id}/role [patch]
func (h *AuthHandler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		responder.Error(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := query.ParseInt64(idStr)
	if id == 0 {
		responder.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req dto.UpdateUserRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.UpdateUserRole(ctx, id, req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 204
// @Router /api/v1/auth/users/{id} [delete]
func (h *AuthHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		responder.Error(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := query.ParseInt64(idStr)
	if id == 0 {
		responder.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := h.service.DeleteUser(ctx, id); err != nil {
		responder.GRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
