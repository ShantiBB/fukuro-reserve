package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils"
)

type AuthHandler struct {
	clients *clients.Clients
}

func NewAuthHandler(clients *clients.Clients) *AuthHandler {
	return &AuthHandler{clients: clients}
}

func authContext(r *http.Request) (context.Context, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, http.ErrNoCookie
	}

	return metadata.NewOutgoingContext(
		r.Context(),
		metadata.Pairs("authorization", authHeader),
	), nil
}

// Register godoc
// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register request"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.Token.RegisterUser(
		r.Context(), &userv1.RegisterUserRequest{
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, tokenResponseFromProto(resp))
}

// Login godoc
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login request"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.Token.LoginUser(
		r.Context(), &userv1.LoginUserRequest{
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, tokenResponseFromLoginProto(resp))
}

// RefreshToken godoc
// @Summary Refresh access token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.Token.RefreshToken(
		r.Context(), &userv1.RefreshTokenRequest{
			RefreshToken: req.RefreshToken,
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, tokenResponseFromRefreshProto(resp))
}

// GetUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} UsersResponse
// @Router /api/v1/auth/users [get]
func (h *AuthHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	page := utils.ParseUint64(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	limit := utils.ParseUint64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 100
	}

	resp, err := h.clients.User.GetUsers(
		ctx, &userv1.GetUsersRequest{
			Page:  page,
			Limit: limit,
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, usersResponseFromProto(resp))
}

// CreateUser godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Create user request"
// @Success 201 {object} UserResponse
// @Router /api/v1/auth/users [post]
func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	var username *string
	if req.Username != "" {
		username = &req.Username
	}

	resp, err := h.clients.User.CreateUser(
		ctx, &userv1.CreateUserRequest{
			Email:    req.Email,
			Username: username,
			Password: req.Password,
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, userResponseFromProto(resp.User))
}

// GetUser godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponse
// @Router /api/v1/auth/users/{id} [get]
func (h *AuthHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := utils.ParseInt64(idStr)
	if id == 0 {
		utils.RespondError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	resp, err := h.clients.User.GetUser(ctx, &userv1.GetUserRequest{Id: id})
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, userResponseFromProto(resp.User))
}

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserRequest true "Update user request"
// @Success 200 {object} UserResponse
// @Router /api/v1/auth/users/{id} [put]
func (h *AuthHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := utils.ParseInt64(idStr)
	if id == 0 {
		utils.RespondError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.User.UpdateUser(
		ctx, &userv1.UpdateUserRequest{
			Id:       id,
			Email:    req.Email,
			Username: req.Username,
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, updateUserResponseFromProto(resp.User))
}

// UpdateUserActivity godoc
// @Summary Update user activity status
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserActivityRequest true "Update activity request"
// @Success 200 {object} UpdateUserActivityResponse
// @Router /api/v1/auth/users/{id}/activity [patch]
func (h *AuthHandler) UpdateUserActivity(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := utils.ParseInt64(idStr)
	if id == 0 {
		utils.RespondError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req UpdateUserActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.User.UpdateUserActivity(
		ctx, &userv1.UpdateUserActivityRequest{
			Id:       id,
			IsActive: wrapperspb.Bool(req.IsActive),
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, UpdateUserActivityResponse{IsActive: resp.IsActive})
}

// UpdateUserRole godoc
// @Summary Update user role
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserRoleRequest true "Update role request"
// @Success 200 {object} UpdateUserRoleResponse
// @Router /api/v1/auth/users/{id}/role [patch]
func (h *AuthHandler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	ctx, err := authContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := utils.ParseInt64(idStr)
	if id == 0 {
		utils.RespondError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req UpdateUserRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.User.UpdateUserRole(
		ctx, &userv1.UpdateUserRoleRequest{
			Id:   id,
			Role: userv1.UserRole(userv1.UserRole_value[req.Role]),
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, UpdateUserRoleResponse{Role: resp.Role.String()})
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
		utils.RespondError(w, http.StatusUnauthorized, "authorization header is required")
		return
	}

	idStr := chi.URLParam(r, "id")
	id := utils.ParseInt64(idStr)
	if id == 0 {
		utils.RespondError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	_, err = h.clients.User.DeleteUser(ctx, &userv1.DeleteUserRequest{Id: id})
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
