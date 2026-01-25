package handler

import (
	"context"
	"errors"
	"net/http"

	helper2 "auth/internal/grpc/lib/utils/helper"
	"auth/internal/http/dto/request"
	"auth/internal/http/dto/response"
	"auth/internal/http/utils/helper"
	"auth/internal/http/utils/validation"
	"auth/pkg/lib/utils/consts"
	"auth/pkg/lib/utils/jwt"
)

const BearerType = "Bearer"

type TokenService interface {
	RegisterByEmail(ctx context.Context, email, password string) (*jwt.Token, error)
	LoginByEmail(ctx context.Context, email, pass string) (*jwt.Token, error)
	RefreshToken(token *jwt.Token) (*jwt.Token, error)
}

// RegisterByEmail    godoc
//
//	@Summary		Register user
//	@Description	Register user by email and return access and refresh tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.InsertUser	true	"User data"
//	@Success		201		{object}	response.Token
//	@Failure		400		{object}	response.ErrorSchema
//	@Failure		409		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Router			/auth/register [post]
func (h *Handler) RegisterByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UserCreate
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	hashPassword, err := helper2.HashPassword(req.Password)
	errHandler := &helper.ErrorHandler{
		BadRequest: consts.ErrPasswordHashing,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	tokens, err := h.svc.RegisterByEmail(ctx, req.Email, hashPassword)
	errHandler = &helper.ErrorHandler{
		Conflict: consts.ErrUniqueUserField,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	tokensResp := response.Token{
		Access:    tokens.Access,
		Refresh:   tokens.Refresh,
		TokenType: BearerType,
	}
	helper.SendSuccess(w, r, http.StatusCreated, tokensResp)
}

// LoginByEmail    godoc
//
//	@Summary		Login user
//	@Description	Login user by email and return access and refresh tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.InsertUser	true	"User data"
//	@Success		200		{object}	response.Token
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Router			/auth/login [post]
func (h *Handler) LoginByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UserCreate
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	tokens, err := h.svc.LoginByEmail(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, consts.ErrInvalidCredentials) || errors.Is(err, consts.ErrUserNotFound) {
			errMsg := response.ErrorResp(consts.ErrInvalidCredentials)
			helper.SendError(w, r, http.StatusUnauthorized, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.ErrInternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(
		w, r, http.StatusOK, response.Token{
			Access:    tokens.Access,
			Refresh:   tokens.Refresh,
			TokenType: BearerType,
		},
	)
}

// RefreshToken    godoc
//
//	@Summary		Refresh access user tokenCreds
//	@Description	Refresh access tokenCreds and return new tokenCreds data
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		object{refresh_token=string}	true	"User data"
//	@Success		200		{object}	response.Token
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Router			/auth/refresh [post]
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req jwt.RefreshToken
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	token := &jwt.Token{Refresh: req.RefreshToken}
	tokens, err := h.svc.RefreshToken(token)
	errHandler := &helper.ErrorHandler{
		Unauthorized: consts.ErrInvalidToken,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	helper.SendSuccess(
		w, r, http.StatusOK, response.Token{
			Access:    tokens.Access,
			Refresh:   tokens.Refresh,
			TokenType: BearerType,
		},
	)
}
