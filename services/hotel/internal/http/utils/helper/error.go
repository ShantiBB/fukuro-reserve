package helper

import (
	"errors"
	"net/http"

	"hotel/internal/http/dto/response"
	"hotel/pkg/utils/consts"
)

type ErrorHandler struct {
	NotFound     error
	Conflict     error
	BadRequest   error
	Unauthorized error
}

func (h *ErrorHandler) Handle(w http.ResponseWriter, r *http.Request, err error) error {
	if err == nil {
		return nil
	}

	if h.NotFound != nil && errors.Is(err, h.NotFound) {
		errMsg := response.ErrorResp(h.NotFound)
		SendError(w, r, http.StatusNotFound, errMsg)
		return err
	}

	if h.Conflict != nil && errors.Is(err, h.Conflict) {
		errMsg := response.ErrorResp(h.Conflict)
		SendError(w, r, http.StatusConflict, errMsg)
		return err
	}

	if h.BadRequest != nil && errors.Is(err, h.BadRequest) {
		errMsg := response.ErrorResp(h.BadRequest)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return err
	}

	if h.Unauthorized != nil && errors.Is(err, h.Unauthorized) {
		errMsg := response.ErrorResp(h.Unauthorized)
		SendError(w, r, http.StatusUnauthorized, errMsg)
		return err
	}

	errMsg := response.ErrorResp(consts.InternalServer)
	SendError(w, r, http.StatusInternalServerError, errMsg)
	return err
}
