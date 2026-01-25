package handler

import (
	"context"
	"net/http"

	helper2 "auth/internal/grpc/lib/utils/helper"
	"auth/internal/http/dto/request"
	"auth/internal/http/dto/response"
	"auth/internal/http/utils/helper"
	"auth/internal/http/utils/pagination"
	"auth/internal/http/utils/validation"
	"auth/internal/repository/models"
	consts2 "auth/pkg/lib/utils/consts"
)

type UserService interface {
	UserCreate(ctx context.Context, user models.CreateUser) (*models.User, error)
	UserGetByID(ctx context.Context, id int64) (*models.User, error)
	UserGetAll(ctx context.Context, page, limit uint64) (*models.UserList, error)
	UserUpdateByID(ctx context.Context, user *models.User) (*models.User, error)
	UserUpdateRoleStatus(ctx context.Context, id int64, role string) error
	UserUpdateActiveStatus(ctx context.Context, id int64, status bool) error
	UserDeleteByID(ctx context.Context, id int64) error
}

// UserCreate    godoc
//
//	@Summary		Create user
//	@Description	Create a new user account from admin provider
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.InsertUser	true	"User data"
//	@Success		201		{object}	response.User
//	@Failure		400		{object}	response.ErrorSchema
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		409		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/users/  [post]
func (h *Handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UserCreate

	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	hashPassword, err := helper2.HashPassword(req.Password)
	errHandler := &helper.ErrorHandler{
		BadRequest: consts2.ErrPasswordHashing,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	newUser := h.UserCreateRequestToEntity(&req, hashPassword)
	createdUser, err := h.svc.UserCreate(ctx, *newUser)
	errHandler = &helper.ErrorHandler{
		Conflict: consts2.ErrUniqueUserField,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	userResponse := h.UserEntityToResponse(createdUser)
	helper.SendSuccess(w, r, http.StatusCreated, userResponse)
}

// UserGetAll    godoc
//
//	@Summary		Get users
//	@Description	Get users from admin or moderator provider
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			page	query		uint	false	"Page"	default(1)
//	@Param			limit	query		uint	false	"Limit"	default(100)
//	@Success		200		{object}	response.UserList
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/users/ [get]
func (h *Handler) UserGetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paginationParams, err := pagination.ParsePaginationQuery(r)
	errHandler := &helper.ErrorHandler{
		BadRequest: consts2.ErrInvalidQueryParam,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	userList, err := h.svc.UserGetAll(ctx, paginationParams.Page, paginationParams.Limit)
	errHandler = &helper.ErrorHandler{}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	users := make([]response.UserShort, 0, len(userList.Users))
	for _, user := range userList.Users {
		userResponse := h.UserShortEntityToResponse(&user)
		users = append(users, *userResponse)
	}

	totalPageCount := (userList.TotalCount + paginationParams.Limit - 1) / paginationParams.Limit
	pageLinks := pagination.BuildPaginationLinks(r, paginationParams, totalPageCount)
	usersResp := response.UserList{
		Users:           users,
		CurrentPage:     paginationParams.Page,
		Limit:           paginationParams.Limit,
		Links:           pageLinks,
		TotalPageCount:  totalPageCount,
		TotalUsersCount: userList.TotalCount,
	}

	helper.SendSuccess(w, r, http.StatusOK, usersResp)
}

// UserGetByID    godoc
//
//	@Summary		Get user by ID
//	@Description	Get user by ID from admin, moderator or owner provider
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	response.User
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/users/{id} [get]
func (h *Handler) UserGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := helper.ParseID(w, r)
	if id == 0 {
		return
	}

	user, err := h.svc.UserGetByID(ctx, id)
	errHandler := &helper.ErrorHandler{
		NotFound: consts2.ErrUserNotFound,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	userResponse := h.UserEntityToResponse(user)
	helper.SendSuccess(w, r, http.StatusOK, userResponse)
}

// UserUpdateByID    godoc
//
//	@Summary		Update user by ID
//	@Description	Update user by ID from admin or owner provider
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"User ID"
//	@Param			request	body		request.UpdateUser	true	"User data"
//	@Success		200		{object}	response.User
//	@Failure		400		{object}	response.ErrorSchema
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		404		{object}	response.ErrorSchema
//	@Failure		409		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/users/{id} [put]
func (h *Handler) UserUpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := helper.ParseID(w, r)
	if id == 0 {
		return
	}

	var req request.UserUpdate
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	user := h.UserUpdateRequestToEntity(&req, id)
	userToUpdate, err := h.svc.UserUpdateByID(ctx, user)
	errHandler := &helper.ErrorHandler{
		NotFound: consts2.ErrUserNotFound,
		Conflict: consts2.ErrUniqueUserField,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	userResponse := h.UserEntityToResponse(userToUpdate)
	helper.SendSuccess(w, r, http.StatusOK, userResponse)
}

// UserUpdateRoleStatus    godoc
//
//	@Summary		Update user role by ID
//	@Description	Update user role by ID from admin or owner provider
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"User ID"
//	@Param			request	body		request.UserRoleStatus	true	"User data"
//	@Success		200		{object}	response.UserRoleStatus
//	@Failure		400		{object}	response.ErrorSchema
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		404		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/users/{id}/role [put]
func (h *Handler) UserUpdateRoleStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := helper.ParseID(w, r)
	if id == 0 {
		return
	}

	var req request.UserRoleStatus
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	err := h.svc.UserUpdateRoleStatus(ctx, id, req.Role)
	errHandler := &helper.ErrorHandler{
		NotFound: consts2.ErrUserNotFound,
		Conflict: consts2.ErrInvalidRole,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	roleStatus := response.UserRoleStatus{
		Message: consts2.UserRoleUpdateSuccess,
		Role:    req.Role,
	}
	helper.SendSuccess(w, r, http.StatusOK, roleStatus)
}

// UserUpdateActiveStatus    godoc
//
//	@Summary		Update user active status by ID
//	@Description	Update user active status by ID from admin or owner provider
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"User ID"
//	@Param			request	body		request.UserActiveStatus	true	"User data"
//	@Success		200		{object}	response.UserActiveStatus
//	@Failure		400		{object}	response.ErrorSchema
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		404		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/users/{id}/status [put]
func (h *Handler) UserUpdateActiveStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := helper.ParseID(w, r)
	if id == 0 {
		return
	}

	var req request.UserActiveStatus
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	err := h.svc.UserUpdateActiveStatus(ctx, id, *req.IsActive)
	errHandler := &helper.ErrorHandler{
		NotFound: consts2.ErrUserNotFound,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	activeStatus := response.UserActiveStatus{
		Message:  consts2.UserActiveUpdateSuccess,
		IsActive: *req.IsActive,
	}
	helper.SendSuccess(w, r, http.StatusOK, activeStatus)
}

// UserDeleteByID    godoc
//
//	@Summary		Delete user by ID
//	@Description	Delete user by ID from admin or owner provider
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/users/{id} [delete]
func (h *Handler) UserDeleteByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := helper.ParseID(w, r)
	if id == 0 {
		return
	}

	err := h.svc.UserDeleteByID(ctx, id)
	errHandler := &helper.ErrorHandler{
		NotFound: consts2.ErrUserNotFound,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
