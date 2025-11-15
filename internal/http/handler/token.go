package handler

import (
	"context"
	"errors"
	"net/http"

	"auth_service/internal/config"
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

	tokens, err := h.svc.RegisterByEmail(ctx, req.Email, hashPassword, h.cfg)
	if err != nil {
		if errors.Is(err, errs.UniqueUserField) {
			errMsg := response.ErrorResp(errs.UniqueUserField)
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

	tokens, err := h.svc.LoginByEmail(ctx, req.Email, req.Password, h.cfg)
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
// @Summary      Refresh access user token
// @Description  Refresh access token and return new token data
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request     body      request.RefreshToken  true  "User data"
// @Success      200         {object}  response.Token
// @Failure      401         {object}  response.Error
// @Failure      500         {object}  response.Error
// @Router       /auth/refresh [post]
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req request.RefreshToken
	if ok := helper.ParseJSON(w, r, &req); !ok {
		return
	}

	token := &jwt.Token{Refresh: req.RefreshToken}
	tokens, err := h.svc.RefreshToken(token, h.cfg)
	if err != nil {
		if errors.Is(err, errs.InvalidToken) {
			errMsg := response.ErrorResp(errs.InvalidToken)
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
