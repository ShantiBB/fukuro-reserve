package handler

import (
	"context"
	"errors"
	"net/http"

	"auth/internal/http/dto/request"
	"auth/internal/http/dto/response"
	"auth/internal/http/lib/helper"
	"fukuro-reserve/pkg/utils/errs"
	"fukuro-reserve/pkg/utils/jwt"
	"fukuro-reserve/pkg/utils/password"
)

const BearerType = "Bearer"

type TokenService interface {
	RegisterByEmail(ctx context.Context, email, password string) (*jwt.Token, error)
	LoginByEmail(ctx context.Context, email, pass string) (*jwt.Token, error)
	RefreshToken(token *jwt.Token) (*jwt.Token, error)
}

// RegisterByEmail    godoc
// @Summary      Register user
// @Description  Register user by email and return access and refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request     body      request.Register      true  "User data"
// @Success      201         {object}  response.Token
// @Failure      400         {object}  response.Error
// @Failure      409         {object}  response.Error
// @Failure      500         {object}  response.Error
// @Router       /auth/register [post]
func (h *Handler) RegisterByEmail(w http.ResponseWriter, r *http.Request) {
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

	tokens, err := h.svc.RegisterByEmail(ctx, req.Email, hashPassword)
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

	tokensResp := response.Token{
		Access:    tokens.Access,
		Refresh:   tokens.Refresh,
		TokenType: BearerType,
	}
	helper.SendSuccess(w, r, http.StatusCreated, tokensResp)
}

// LoginByEmail    godoc
// @Summary      Login user
// @Description  Login user by email and return access and refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request     body      request.LoginByEmail  true  "User data"
// @Success      200         {object}  response.Token
// @Failure      401         {object}  response.Error
// @Failure      500         {object}  response.Error
// @Router       /auth/login [post]
func (h *Handler) LoginByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UserCreate
	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	tokens, err := h.svc.LoginByEmail(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, errs.InvalidCredentials) || errors.Is(err, errs.UserNotFound) {
			errMsg := response.ErrorResp(errs.InvalidCredentials)
			helper.SendError(w, r, http.StatusUnauthorized, errMsg)
			return
		}
		errMsg := response.ErrorResp(errs.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(w, r, http.StatusOK, response.Token{
		Access:    tokens.Access,
		Refresh:   tokens.Refresh,
		TokenType: BearerType,
	})
}

// RefreshToken    godoc
// @Summary      Refresh access user tokenCreds
// @Description  Refresh access tokenCreds and return new tokenCreds data
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request     body      request.RefreshToken  true  "User data"
// @Success      200         {object}  response.Token
// @Failure      401         {object}  response.Error
// @Failure      500         {object}  response.Error
// @Router       /auth/refresh [post]
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req jwt.RefreshToken
	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	token := &jwt.Token{Refresh: req.RefreshToken}
	tokens, err := h.svc.RefreshToken(token)
	if err != nil {
		if errors.Is(err, errs.InvalidRefreshToken) {
			errMsg := response.ErrorResp(errs.InvalidRefreshToken)
			helper.SendError(w, r, http.StatusUnauthorized, errMsg)
			return
		}
		errMsg := response.ErrorResp(errs.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(w, r, http.StatusOK, response.Token{
		Access:    tokens.Access,
		Refresh:   tokens.Refresh,
		TokenType: BearerType,
	})
}
