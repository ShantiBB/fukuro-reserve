package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"auth/internal/http/dto/request"
	"auth/internal/http/dto/response"
	"auth/internal/http/lib/helper"
	"auth/internal/repository/models"
	"fukuro-reserve/pkg/utils/errs"
	"fukuro-reserve/pkg/utils/password"
)

type UserService interface {
	UserCreate(ctx context.Context, user models.UserCreate) (*models.User, error)
	UserGetByID(ctx context.Context, id int64) (*models.User, error)
	UserList(ctx context.Context) ([]models.User, error)
	UserUpdateByID(ctx context.Context, user *models.User) (*models.User, error)
	UserDeleteByID(ctx context.Context, id int64) error
}

// UserCreate    godoc
// @Summary      Create user
// @Description  Create a new user account from admin provider
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      request.UserCreate  true  "User data"
// @Success      201      {object}  response.User
// @Failure      400      {object}  response.Error
// @Failure      401      {object}  response.Error
// @Failure      409      {object}  response.Error
// @Failure      500      {object}  response.Error
// @Security     Bearer
// @Router       /users/  [post]
func (h *Handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UserCreate

	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		errMsg := response.ErrorResp(errs.PasswordHashing)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	newUser := h.UserCreateRequestToEntity(&req, hashPassword)
	createdUser, err := h.svc.UserCreate(ctx, *newUser)
	if err != nil {
		if errors.Is(err, errs.UniqueEmailField) {
			errMsg := response.ErrorResp(errs.UniqueEmailField)
			helper.SendError(w, r, http.StatusConflict, errMsg)
			return
		}
		errMsg := response.ErrorResp(errs.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	userResponse := h.UserEntityToResponse(createdUser)
	helper.SendSuccess(w, r, http.StatusCreated, userResponse)
}

// UserList    godoc
// @Summary      Get users
// @Description  Get users from admin or moderator provider
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.User
// @Failure      401  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Security     Bearer
// @Router       /users/ [get]
func (h *Handler) UserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.svc.UserList(ctx)
	if err != nil {
		errMsg := response.ErrorResp(errs.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	usersResp := make([]response.User, 0, len(users))
	for _, user := range users {
		userResponse := h.UserEntityToResponse(&user)
		usersResp = append(usersResp, *userResponse)
	}
	helper.SendSuccess(w, r, http.StatusOK, usersResp)
}

// UserGetByID    godoc
// @Summary      Get user by ID
// @Description  Get user by ID from admin, moderator or owner provider
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int true  "User ID"
// @Success      200  {object}  response.User
// @Failure      400  {object}  response.Error
// @Failure      401  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Security     Bearer
// @Router       /users/{id} [get]
func (h *Handler) UserGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := response.ErrorResp(errs.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	user, err := h.svc.UserGetByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.UserNotFound) {
			errMsg := response.ErrorResp(errs.UserNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(errs.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	userResponse := h.UserEntityToResponse(user)
	helper.SendSuccess(w, r, http.StatusOK, userResponse)
}

// UserUpdateByID    godoc
// @Summary      Update user by ID
// @Description  Update user by ID from admin or owner provider
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id          path      int true  "User ID"
// @Param        request     body      request.UserUpdate  true  "User data"
// @Success      200         {object}  response.User
// @Failure      400         {object}  response.Error
// @Failure      401         {object}  response.Error
// @Failure      404         {object}  response.Error
// @Failure      409         {object}  response.Error
// @Failure      500         {object}  response.Error
// @Security     Bearer
// @Router       /users/{id} [put]
func (h *Handler) UserUpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := response.ErrorResp(errs.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	var req request.UserUpdate
	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	user := h.UserUpdateRequestToEntity(&req, id)
	userToUpdate, err := h.svc.UserUpdateByID(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, errs.UniqueEmailField):
			errMsg := response.ErrorResp(errs.UniqueEmailField)
			helper.SendError(w, r, http.StatusConflict, errMsg)
			return
		case errors.Is(err, errs.UserNotFound):
			errMsg := response.ErrorResp(errs.UserNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(errs.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	userResponse := h.UserEntityToResponse(userToUpdate)

	helper.SendSuccess(w, r, http.StatusOK, userResponse)
}

// UserDeleteByID    godoc
// @Summary      Delete user by ID
// @Description  Delete user by ID from admin or owner provider
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int true  "User ID"
// @Success      204  {object}  nil
// @Failure      400  {object}  response.Error
// @Failure      401  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Security     Bearer
// @Router       /users/{id} [delete]
func (h *Handler) UserDeleteByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := response.ErrorResp(errs.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	if err = h.svc.UserDeleteByID(ctx, id); err != nil {
		if errors.Is(err, errs.UserNotFound) {
			errMsg := response.ErrorResp(errs.UserNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(errs.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
