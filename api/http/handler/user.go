package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"auth_service/api/http/lib/validation"
	"auth_service/api/http/schemas"
	"auth_service/package/utils/password"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	// Update(w http.ResponseWriter, r *http.Request)
	// Delete(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req schemas.UserCreateRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.sendError(w, r, http.StatusBadRequest, "Bad request")
		return
	}

	if errResp := validation.CheckErrors(&req); errResp != nil {
		h.sendError(w, r, http.StatusBadRequest, errResp)
		return
	}

	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, "Error hashing password")
		return
	}

	newUser := h.UserCreateRequestToEntity(&req, hashPassword)
	createdUser, err := h.UserService.Create(ctx, *newUser)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	userResponse := h.UserEntityToResponse(createdUser)
	h.sendJSON(w, r, http.StatusCreated, userResponse)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.UserService.Get(ctx, id)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	userResponse := h.UserEntityToResponse(user)
	h.sendJSON(w, r, http.StatusOK, userResponse)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.UserService.GetAll(ctx)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	usersResp := make([]schemas.UserResponse, 0, len(users))
	for _, user := range users {
		userResponse := h.UserEntityToResponse(&user)
		usersResp = append(usersResp, *userResponse)
	}
	h.sendJSON(w, r, http.StatusOK, usersResp)
}
