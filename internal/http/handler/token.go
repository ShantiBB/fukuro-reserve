package handler

import (
	"context"
	"errors"
	"net/http"

	"auth_service/internal/config"
	"auth_service/internal/http/lib/schemas"
	"auth_service/internal/http/lib/schemas/request"
	"auth_service/internal/http/lib/schemas/response"
	"auth_service/package/utils/errs"
	"auth_service/package/utils/helper"
	"auth_service/package/utils/jwt"
	"auth_service/package/utils/password"
)

const BearerType = "Bearer"

type TokenService interface {
	RegisterByEmail(ctx context.Context, email, password string, cfg *config.Config) (*jwt.Token, error)
	LoginByEmail(ctx context.Context, email, pass string, cfg *config.Config) (*jwt.Token, error)
	RefreshToken(token *jwt.Token, cfg *config.Config) (*jwt.Token, error)
}

func (h *Handler) RegisterByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.Register
	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Error hashing password")
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	tokens, err := h.svc.RegisterByEmail(ctx, req.Email, hashPassword, h.cfg)
	if err != nil {
		if errors.Is(err, errs.UniqueUserField) {
			errMsg := schemas.NewErrorResponse("Email or username already exists")
			helper.SendError(w, r, http.StatusConflict, errMsg)
			return
		}
		errMsg := schemas.NewErrorResponse("Error registering user")
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	tokensResp := response.Token{
		Access:    tokens.Access,
		Refresh:   tokens.Refresh,
		TokenType: BearerType,
	}
	helper.SendSuccess(w, r, http.StatusCreated, tokensResp)
}

func (h *Handler) LoginByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.LoginByEmail
	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	tokens, err := h.svc.LoginByEmail(ctx, req.Email, req.Password, h.cfg)
	if err != nil {
		if errors.Is(err, errs.InvalidCredentials) || errors.Is(err, errs.UserNotFound) {
			errMsg := schemas.NewErrorResponse("Invalid email or password")
			helper.SendError(w, r, http.StatusUnauthorized, errMsg)
			return
		}
		errMsg := schemas.NewErrorResponse("Error logging in user")
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(w, r, http.StatusOK, response.Token{
		Access:    tokens.Access,
		Refresh:   tokens.Refresh,
		TokenType: BearerType,
	})
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req request.RefreshToken
	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	token := &jwt.Token{Refresh: req.RefreshToken}
	tokens, err := h.svc.RefreshToken(token, h.cfg)
	if err != nil {
		if errors.Is(err, errs.InvalidToken) {
			errMsg := schemas.NewErrorResponse(errs.InvalidToken.Error())
			helper.SendError(w, r, http.StatusUnauthorized, errMsg)
			return
		}
		errMsg := schemas.NewErrorResponse("Error refreshing token")
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(w, r, http.StatusOK, response.Token{
		Access:    tokens.Access,
		Refresh:   tokens.Refresh,
		TokenType: BearerType,
	})
}
